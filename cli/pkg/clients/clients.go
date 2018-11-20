package clients

import (
	"fmt"
	"strings"

	glooV1 "github.com/solo-io/supergloo/pkg/api/external/gloo/v1"

	"github.com/solo-io/solo-kit/pkg/errors"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	cliConstants "github.com/solo-io/supergloo/cli/pkg/constants"
	"github.com/solo-io/supergloo/cli/pkg/model/info"
	sgConstants "github.com/solo-io/supergloo/pkg/constants"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube"
	"github.com/solo-io/solo-kit/pkg/utils/kubeutils"
	superglooV1 "github.com/solo-io/supergloo/pkg/api/v1"
	k8sApiExt "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/typed/apiextensions/v1beta1"
	k8s "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

type SuperglooInfoClient interface {
	ListResourceTypes() ([]string, error)
	ListResources(resourceType, resourceName string) (info.ResourceInfo, error)
}

type KubernetesSuperglooClient struct {
	kubeCrdClient     *k8sApiExt.CustomResourceDefinitionInterface
	meshClient        *superglooV1.MeshClient
	routingRuleClient *superglooV1.RoutingRuleClient
	upstreamClient    *glooV1.UpstreamClient
}

func NewClient() (SuperglooInfoClient, error) {
	config, err := kubeutils.GetConfig("", "")
	if err != nil {
		return nil, fmt.Errorf(cliConstants.KubeConfigError, err)
	}

	crdClient, err := getCrdClient(config)
	if err != nil {
		return nil, err
	}

	meshClient, err := superglooV1.NewMeshClient(&factory.KubeResourceClientFactory{
		Crd:         superglooV1.MeshCrd,
		Cfg:         config,
		SharedCache: kube.NewKubeCache(),
	})
	if err != nil {
		return nil, err
	}
	if err = meshClient.Register(); err != nil {
		return nil, err
	}

	routingRuleClient, err := superglooV1.NewRoutingRuleClient(&factory.KubeResourceClientFactory{
		Crd:         superglooV1.RoutingRuleCrd,
		Cfg:         config,
		SharedCache: kube.NewKubeCache(),
	})
	if err != nil {
		return nil, err
	}
	if err = routingRuleClient.Register(); err != nil {
		return nil, err
	}

	upstreamClient, err := glooV1.NewUpstreamClient(&factory.KubeResourceClientFactory{
		Crd:         glooV1.UpstreamCrd,
		Cfg:         config,
		SharedCache: kube.NewKubeCache(),
	})
	if err != nil {
		return nil, err
	}
	if err = upstreamClient.Register(); err != nil {
		return nil, err
	}

	client := &KubernetesSuperglooClient{
		kubeCrdClient:     crdClient,
		meshClient:        &meshClient,
		routingRuleClient: &routingRuleClient,
		upstreamClient:    &upstreamClient,
	}

	return client, nil
}

// Get a list of all supergloo resources
func (client *KubernetesSuperglooClient) ListResourceTypes() ([]string, error) {
	crdList, err := (*client.kubeCrdClient).List(k8s.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("Error retrieving supergloo resource types. Cause: %v \n", err)
	}

	superglooCRDs := make([]string, 0)
	for _, crd := range crdList.Items {
		if strings.Contains(crd.Name, cliConstants.SuperglooGroupName) {
			parts := strings.Split(crd.Name, ".")
			if len(parts) > 0 {
				superglooCRDs = append(superglooCRDs, parts[0])
			}
		}
	}
	return superglooCRDs, nil
}

func (client *KubernetesSuperglooClient) ListResources(resourceType, resourceName string) (info.ResourceInfo, error) {
	// TODO(marco): make code more generic. Ideally we don't want to enumerate the different options, but I could not
	// find an interface that all of the generated clients implement
	switch resourceType {
	case "meshes":
		if resourceName == "" {
			meshList, err := (*client.meshClient).List(sgConstants.SuperglooNamespace, clients.ListOpts{})
			if err != nil {
				return nil, err
			}
			return info.FromList(&meshList), nil
		} else {
			mesh, err := (*client.meshClient).Read(sgConstants.SuperglooNamespace, resourceName, clients.ReadOpts{})
			if err != nil {
				return nil, err
			}
			return info.From(mesh), nil
		}
	default:
		// Should not happen since we validate the resource
		return nil, errors.Errorf(cliConstants.UnknownResourceTypeMsg, resourceType)
	}
}

// Return a client to query kubernetes CRDs
func getCrdClient(config *rest.Config) (*k8sApiExt.CustomResourceDefinitionInterface, error) {
	apiExtClient, err := k8sApiExt.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("Error building Kubernetes client: %v \n", err)
	}
	crdClient := apiExtClient.CustomResourceDefinitions()
	return &crdClient, nil
}
