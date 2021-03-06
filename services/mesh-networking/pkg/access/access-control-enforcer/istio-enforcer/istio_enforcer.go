package istio_enforcer

import (
	"context"

	discovery_v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1"
	"github.com/solo-io/service-mesh-hub/pkg/clients/istio/security"
	"github.com/solo-io/service-mesh-hub/services/common/constants"
	mc_manager "github.com/solo-io/service-mesh-hub/services/common/multicluster/manager"
	global_access_control_enforcer "github.com/solo-io/service-mesh-hub/services/mesh-networking/pkg/access/access-control-enforcer"
	istio_federation "github.com/solo-io/service-mesh-hub/services/mesh-networking/pkg/federation/resolver/meshes/istio"
	security_v1beta1 "istio.io/api/security/v1beta1"
	"istio.io/api/type/v1beta1"
	client_security_v1beta1 "istio.io/client-go/pkg/apis/security/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	EnforcerId                        = "istio_enforcer"
	GlobalAccessControlAuthPolicyName = "global-access-control"
	IngressGatewayAuthPolicy          = "ingress-policy"
)

type istioEnforcer struct {
	dynamicClientGetter     mc_manager.DynamicClientGetter
	authPolicyClientFactory security.AuthorizationPolicyClientFactory
}

type IstioEnforcer global_access_control_enforcer.AccessPolicyMeshEnforcer

func NewIstioEnforcer(
	dynamicClientGetter mc_manager.DynamicClientGetter,
	authPolicyClientFactory security.AuthorizationPolicyClientFactory,
) IstioEnforcer {
	return &istioEnforcer{
		authPolicyClientFactory: authPolicyClientFactory,
		dynamicClientGetter:     dynamicClientGetter,
	}
}

func (i *istioEnforcer) Name() string {
	return EnforcerId
}

func (i *istioEnforcer) StartEnforcing(ctx context.Context, meshes []*discovery_v1alpha1.Mesh) error {
	for _, mesh := range meshes {
		if mesh.Spec.GetIstio() == nil {
			continue
		}
		clientForCluster, err := i.dynamicClientGetter.GetClientForCluster(mesh.Spec.GetCluster().GetName())
		if err != nil {
			return err
		}
		authPolicyClient := i.authPolicyClientFactory(clientForCluster)
		if err := i.ensureIngressGatewayPolicy(ctx, mesh, authPolicyClient); err != nil {
			return err
		}
		if err := i.ensureGlobalAuthPolicy(ctx, mesh, authPolicyClient); err != nil {
			return err
		}
	}
	return nil
}

func (i *istioEnforcer) StopEnforcing(ctx context.Context, meshes []*discovery_v1alpha1.Mesh) error {
	for _, mesh := range meshes {
		if mesh.Spec.GetIstio() == nil {
			continue
		}
		if err := i.stopEnforcingForMesh(ctx, mesh); err != nil {
			return err
		}
	}
	return nil
}

func (i *istioEnforcer) ensureGlobalAuthPolicy(
	ctx context.Context,
	mesh *discovery_v1alpha1.Mesh,
	authPolicyClient security.AuthorizationPolicyClient,
) error {
	// The following config denies all traffic in the mesh because it defaults to an ALLOW rule that doesn't match any requests,
	// thereby denying traffic unless explicitly allowed by the user through additional AuthorizationPolicies.
	// https://istio.io/docs/reference/config/security/authorization-policy/#AuthorizationPolicy
	globalAccessControlAuthPolicy := &client_security_v1beta1.AuthorizationPolicy{
		ObjectMeta: v1.ObjectMeta{
			Name:      GlobalAccessControlAuthPolicyName,
			Namespace: mesh.Spec.GetIstio().GetInstallation().GetInstallationNamespace(),
			Labels:    constants.OwnedBySMHLabel,
		},
		Spec: security_v1beta1.AuthorizationPolicy{},
	}
	return authPolicyClient.UpsertSpec(ctx, globalAccessControlAuthPolicy)
}

func (i *istioEnforcer) ensureIngressGatewayPolicy(
	ctx context.Context,
	mesh *discovery_v1alpha1.Mesh,
	authPolicyClient security.AuthorizationPolicyClient,
) error {
	// The following config allows all traffic into the "istio-ingressgateway", which in Service Mesh Hub is
	// the gateway used for multi cluster traffic. Authorization is then handled by the invidual workloads which traffic
	// is forwarded to.
	ingressGatewayAllowAllPolicy := &client_security_v1beta1.AuthorizationPolicy{
		ObjectMeta: v1.ObjectMeta{
			Name:      IngressGatewayAuthPolicy,
			Namespace: mesh.Spec.GetIstio().GetInstallation().GetInstallationNamespace(),
			Labels:    constants.OwnedBySMHLabel,
		},
		Spec: security_v1beta1.AuthorizationPolicy{
			Action: security_v1beta1.AuthorizationPolicy_ALLOW,
			// According to the Istio docs on AuthorizationPolicy a single empty rule allows all traffic
			// https://istio.io/docs/reference/config/security/authorization-policy/#AuthorizationPolicy
			Rules: []*security_v1beta1.Rule{{}},
			Selector: &v1beta1.WorkloadSelector{
				MatchLabels: istio_federation.BuildGatewayWorkloadSelector(),
			},
		},
	}
	return authPolicyClient.UpsertSpec(ctx, ingressGatewayAllowAllPolicy)
}

func (i *istioEnforcer) stopEnforcingForMesh(
	ctx context.Context,
	mesh *discovery_v1alpha1.Mesh,
) error {
	clientForCluster, err := i.dynamicClientGetter.GetClientForCluster(mesh.Spec.GetCluster().GetName())
	if err != nil {
		return err
	}
	authPolicyClient := i.authPolicyClientFactory(clientForCluster)
	globalAuthPolicyKey := client.ObjectKey{
		Name:      GlobalAccessControlAuthPolicyName,
		Namespace: mesh.Spec.GetIstio().GetInstallation().GetInstallationNamespace(),
	}
	if err = i.deleteIfExists(ctx, globalAuthPolicyKey, authPolicyClient); err != nil {
		return err
	}
	gatewayAuthPolicyKey := client.ObjectKey{
		Name:      IngressGatewayAuthPolicy,
		Namespace: mesh.Spec.GetIstio().GetInstallation().GetInstallationNamespace(),
	}
	return i.deleteIfExists(ctx, gatewayAuthPolicyKey, authPolicyClient)
}

func (i *istioEnforcer) deleteIfExists(
	ctx context.Context,
	objKey client.ObjectKey,
	policyClient security.AuthorizationPolicyClient,
) error {
	_, err := policyClient.Get(ctx, objKey)
	if err != nil {
		// If it cannot be found, do not attempt to delete, and return no error
		if errors.IsNotFound(err) {
			return nil
		}
		return err
	} else {
		return policyClient.Delete(ctx, objKey)
	}
}
