
---
title: "traffic_policy.proto"
---

## Package : `networking.zephyr.solo.io`



<a name="top"></a>

<a name="API Reference for traffic_policy.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## traffic_policy.proto


## Table of Contents
  - [TrafficPolicySpec](#networking.zephyr.solo.io.TrafficPolicySpec)
  - [TrafficPolicySpec.CorsPolicy](#networking.zephyr.solo.io.TrafficPolicySpec.CorsPolicy)
  - [TrafficPolicySpec.FaultInjection](#networking.zephyr.solo.io.TrafficPolicySpec.FaultInjection)
  - [TrafficPolicySpec.FaultInjection.Abort](#networking.zephyr.solo.io.TrafficPolicySpec.FaultInjection.Abort)
  - [TrafficPolicySpec.FaultInjection.Delay](#networking.zephyr.solo.io.TrafficPolicySpec.FaultInjection.Delay)
  - [TrafficPolicySpec.HeaderManipulation](#networking.zephyr.solo.io.TrafficPolicySpec.HeaderManipulation)
  - [TrafficPolicySpec.HeaderManipulation.AppendRequestHeadersEntry](#networking.zephyr.solo.io.TrafficPolicySpec.HeaderManipulation.AppendRequestHeadersEntry)
  - [TrafficPolicySpec.HeaderManipulation.AppendResponseHeadersEntry](#networking.zephyr.solo.io.TrafficPolicySpec.HeaderManipulation.AppendResponseHeadersEntry)
  - [TrafficPolicySpec.HeaderMatcher](#networking.zephyr.solo.io.TrafficPolicySpec.HeaderMatcher)
  - [TrafficPolicySpec.HttpMatcher](#networking.zephyr.solo.io.TrafficPolicySpec.HttpMatcher)
  - [TrafficPolicySpec.HttpMethod](#networking.zephyr.solo.io.TrafficPolicySpec.HttpMethod)
  - [TrafficPolicySpec.Mirror](#networking.zephyr.solo.io.TrafficPolicySpec.Mirror)
  - [TrafficPolicySpec.MultiDestination](#networking.zephyr.solo.io.TrafficPolicySpec.MultiDestination)
  - [TrafficPolicySpec.MultiDestination.WeightedDestination](#networking.zephyr.solo.io.TrafficPolicySpec.MultiDestination.WeightedDestination)
  - [TrafficPolicySpec.MultiDestination.WeightedDestination.SubsetEntry](#networking.zephyr.solo.io.TrafficPolicySpec.MultiDestination.WeightedDestination.SubsetEntry)
  - [TrafficPolicySpec.QueryParameterMatcher](#networking.zephyr.solo.io.TrafficPolicySpec.QueryParameterMatcher)
  - [TrafficPolicySpec.RetryPolicy](#networking.zephyr.solo.io.TrafficPolicySpec.RetryPolicy)
  - [TrafficPolicySpec.StringMatch](#networking.zephyr.solo.io.TrafficPolicySpec.StringMatch)
  - [TrafficPolicyStatus](#networking.zephyr.solo.io.TrafficPolicyStatus)
  - [TrafficPolicyStatus.TranslatorError](#networking.zephyr.solo.io.TrafficPolicyStatus.TranslatorError)







<a name="networking.zephyr.solo.io.TrafficPolicySpec"></a>

### TrafficPolicySpec
A routing rule applies some L7 routing features to an existing mesh. Routing rules specify the following for all requests: - originating from from **source pods** - sent to **destination services** - matching one or more **request matcher** apply the specified TrafficPolicySpec the routing configuration that will be applied to the mesh(es)<br>Throughout the documentation below, the term "destination" or "destination service" refers to the underlying Kubernetes service that is represented in Service Mesh Hub as a MeshService.<br>NB: If any additional TrafficPolicy action fields (i.e. non selection related fields) are added, the TrafficPolicy Merger's "AreTrafficPolicyActionsEqual" method must be updated to reflect the new field.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sourceSelector | [core.zephyr.solo.io.WorkloadSelector](#core.zephyr.solo.io.WorkloadSelector) |  | Requests originating from these pods will have the rule applied. Leave empty to have all pods in the mesh apply these rules.<br>Note: Source Selectors are ignored when TrafficPolicies are applied to pods in a Linkerd mesh. TrafficPolicies will apply to all selected destinations in Linkerd, regardless of the source.<br>Note: If using the ServiceSelector.Matcher, specifying clusters is currently not supported in Istio. |
| destinationSelector | [core.zephyr.solo.io.ServiceSelector](#core.zephyr.solo.io.ServiceSelector) |  | Requests destined for these k8s services will have the rule applied. Leave empty to apply to all destination k8s services in the mesh. |
| httpRequestMatchers | [][TrafficPolicySpec.HttpMatcher](#networking.zephyr.solo.io.TrafficPolicySpec.HttpMatcher) | repeated | If specified, this rule will only apply to http requests matching these conditions. Within a single matcher, all conditions must be satisfied for a match to occur. Between matchers, at least one matcher must be satisfied for the TrafficPolicy to apply. NB: Linkerd only supports matching on Request Path and Method |
| trafficShift | [TrafficPolicySpec.MultiDestination](#networking.zephyr.solo.io.TrafficPolicySpec.MultiDestination) |  | A routing rule can have one of several types. Note: types imported from Istio will be replaced with our own simpler types, this is just a place to start from.<br>Enables traffic shifting, i.e. to reroute requests to a different service, to a subset of pods based on their label, and/or split traffic between multiple services. |
| faultInjection | [TrafficPolicySpec.FaultInjection](#networking.zephyr.solo.io.TrafficPolicySpec.FaultInjection) |  | Enable fault injection on requests. |
| requestTimeout | [google.protobuf.Duration](#google.protobuf.Duration) |  | Set a timeout on requests. |
| retries | [TrafficPolicySpec.RetryPolicy](#networking.zephyr.solo.io.TrafficPolicySpec.RetryPolicy) |  | Set a retry policy on requests. |
| corsPolicy | [TrafficPolicySpec.CorsPolicy](#networking.zephyr.solo.io.TrafficPolicySpec.CorsPolicy) |  | Set a Cross-Origin Resource Sharing policy (CORS) for requests. Refer to https://developer.mozilla.org/en-US/docs/Web/HTTP/Access_control_CORS for further details about cross origin resource sharing. |
| mirror | [TrafficPolicySpec.Mirror](#networking.zephyr.solo.io.TrafficPolicySpec.Mirror) |  | Mirror HTTP traffic to a another destination. Traffic will still be sent to its original destination as normal. |
| headerManipulation | [TrafficPolicySpec.HeaderManipulation](#networking.zephyr.solo.io.TrafficPolicySpec.HeaderManipulation) |  | Manipulate request and response headers. |






<a name="networking.zephyr.solo.io.TrafficPolicySpec.CorsPolicy"></a>

### TrafficPolicySpec.CorsPolicy



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| allowOrigins | [][TrafficPolicySpec.StringMatch](#networking.zephyr.solo.io.TrafficPolicySpec.StringMatch) | repeated | String patterns that match allowed origins. An origin is allowed if any of the string matchers match. If a match is found, then the outgoing Access-Control-Allow-Origin would be set to the origin as provided by the client. |
| allowMethods | [][string](#string) | repeated | List of HTTP methods allowed to access the resource. The content will be serialized into the Access-Control-Allow-Methods header. |
| allowHeaders | [][string](#string) | repeated | List of HTTP headers that can be used when requesting the resource. Serialized to Access-Control-Allow-Headers header. |
| exposeHeaders | [][string](#string) | repeated | A white list of HTTP headers that the browsers are allowed to access. Serialized into Access-Control-Expose-Headers header. |
| maxAge | [google.protobuf.Duration](#google.protobuf.Duration) |  | Specifies how long the results of a preflight request can be cached. Translates to the `Access-Control-Max-Age` header. |
| allowCredentials | [google.protobuf.BoolValue](#google.protobuf.BoolValue) |  | Indicates whether the caller is allowed to send the actual request (not the preflight) using credentials. Translates to `Access-Control-Allow-Credentials` header. |






<a name="networking.zephyr.solo.io.TrafficPolicySpec.FaultInjection"></a>

### TrafficPolicySpec.FaultInjection
FaultInjection can be used to specify one or more faults to inject while forwarding http requests to the destination specified in a route. Faults include aborting the Http request from downstream service, and/or delaying proxying of requests. A fault rule MUST HAVE delay or abort.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| delay | [TrafficPolicySpec.FaultInjection.Delay](#networking.zephyr.solo.io.TrafficPolicySpec.FaultInjection.Delay) |  | Delay requests before forwarding, emulating various failures such as network issues, overloaded upstream service, etc. |
| abort | [TrafficPolicySpec.FaultInjection.Abort](#networking.zephyr.solo.io.TrafficPolicySpec.FaultInjection.Abort) |  | Abort Http request attempts and return error codes back to downstream service, giving the impression that the upstream service is faulty. |
| percentage | [double](#double) |  | Percentage of requests to be faulted with the error code provided. Values range between 0 and 100 |






<a name="networking.zephyr.solo.io.TrafficPolicySpec.FaultInjection.Abort"></a>

### TrafficPolicySpec.FaultInjection.Abort
The _httpStatus_ field is used to indicate the HTTP status code to return to the caller. The optional _percentage_ field can be used to only abort a certain percentage of requests. If not specified, all requests are aborted.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| httpStatus | [int32](#int32) |  | REQUIRED. HTTP status code to use to abort the Http request. |






<a name="networking.zephyr.solo.io.TrafficPolicySpec.FaultInjection.Delay"></a>

### TrafficPolicySpec.FaultInjection.Delay
The _fixedDelay_ field is used to indicate the amount of delay in seconds. The optional _percentage_ field can be used to only delay a certain percentage of requests. If left unspecified, all request will be delayed.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| fixedDelay | [google.protobuf.Duration](#google.protobuf.Duration) |  | Add a fixed delay before forwarding the request. Format: 1h/1m/1s/1ms. MUST be >=1ms. |
| exponentialDelay | [google.protobuf.Duration](#google.protobuf.Duration) |  | $hide_from_docs |






<a name="networking.zephyr.solo.io.TrafficPolicySpec.HeaderManipulation"></a>

### TrafficPolicySpec.HeaderManipulation
Manipulate request and response headers.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| removeResponseHeaders | [][string](#string) | repeated | HTTP headers to remove before returning a response to the caller. |
| appendResponseHeaders | [][TrafficPolicySpec.HeaderManipulation.AppendResponseHeadersEntry](#networking.zephyr.solo.io.TrafficPolicySpec.HeaderManipulation.AppendResponseHeadersEntry) | repeated | Additional HTTP headers to add before returning a response to the caller. |
| removeRequestHeaders | [][string](#string) | repeated | HTTP headers to remove before forwarding a request to the destination service. |
| appendRequestHeaders | [][TrafficPolicySpec.HeaderManipulation.AppendRequestHeadersEntry](#networking.zephyr.solo.io.TrafficPolicySpec.HeaderManipulation.AppendRequestHeadersEntry) | repeated | Additional HTTP headers to add before forwarding a request to the destination service. |






<a name="networking.zephyr.solo.io.TrafficPolicySpec.HeaderManipulation.AppendRequestHeadersEntry"></a>

### TrafficPolicySpec.HeaderManipulation.AppendRequestHeadersEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="networking.zephyr.solo.io.TrafficPolicySpec.HeaderManipulation.AppendResponseHeadersEntry"></a>

### TrafficPolicySpec.HeaderManipulation.AppendResponseHeadersEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="networking.zephyr.solo.io.TrafficPolicySpec.HeaderMatcher"></a>

### TrafficPolicySpec.HeaderMatcher



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Specifies the name of the header in the request. |
| value | [string](#string) |  | Specifies the value of the header. If the value is absent a request that has the name header will match, regardless of the header’s value. |
| regex | [bool](#bool) |  | Specifies whether the header value should be treated as regex or not. |
| invertMatch | [bool](#bool) |  | If set to true, the result of the match will be inverted. Defaults to false.<br>Examples: name=foo, invert_match=true: matches if no header named `foo` is present name=foo, value=bar, invert_match=true: matches if no header named `foo` with value `bar` is present name=foo, value=``\d{3}``, regex=true, invert_match=true: matches if no header named `foo` with a value consisting of three integers is present |






<a name="networking.zephyr.solo.io.TrafficPolicySpec.HttpMatcher"></a>

### TrafficPolicySpec.HttpMatcher
Parameters for matching routes. All specified conditions must be satisfied for a match to occur.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| prefix | [string](#string) |  | If specified, the route is a prefix rule meaning that the prefix must match the beginning of the *:path* header. |
| exact | [string](#string) |  | If specified, the route is an exact path rule meaning that the path must exactly match the *:path* header once the query string is removed. |
| regex | [string](#string) |  | If specified, the route is a regular expression rule meaning that the regex must match the *:path* header once the query string is removed. The entire path (without the query string) must match the regex. The rule will not match if only a sub-sequence of the *:path* header matches the regex. The regex grammar is defined `here <http://en.cppreference.com/w/cpp/regex/ecmascript>`_. |
| headers | [][TrafficPolicySpec.HeaderMatcher](#networking.zephyr.solo.io.TrafficPolicySpec.HeaderMatcher) | repeated | Specifies a set of headers which requests must match in entirety (all headers must match). |
| queryParameters | [][TrafficPolicySpec.QueryParameterMatcher](#networking.zephyr.solo.io.TrafficPolicySpec.QueryParameterMatcher) | repeated | Specifies a set of URL query parameters which requests must match in entirety (all query params must match). The router will check the query string from the *path* header against all the specified query parameters |
| method | [TrafficPolicySpec.HttpMethod](#networking.zephyr.solo.io.TrafficPolicySpec.HttpMethod) |  | HTTP Method/Verb to match on. If none specified, the matcher will ignore the HTTP Method |






<a name="networking.zephyr.solo.io.TrafficPolicySpec.HttpMethod"></a>

### TrafficPolicySpec.HttpMethod
Express an optional HttpMethod by wrapping it in a nillable message.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| method | [core.zephyr.solo.io.HttpMethodValue](#core.zephyr.solo.io.HttpMethodValue) |  |  |






<a name="networking.zephyr.solo.io.TrafficPolicySpec.Mirror"></a>

### TrafficPolicySpec.Mirror



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| destination | [core.zephyr.solo.io.ResourceRef](#core.zephyr.solo.io.ResourceRef) |  | Destination to mirror traffic to |
| percentage | [double](#double) |  | Percentage of traffic to mirror. If absent, 100% will be mirrored. Values range between 0 and 100 |
| port | [uint32](#uint32) |  | Port on the destination service to receive traffic. If multiple are found, and none are specified, then the configuration will be considered invalid. |






<a name="networking.zephyr.solo.io.TrafficPolicySpec.MultiDestination"></a>

### TrafficPolicySpec.MultiDestination



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| destinations | [][TrafficPolicySpec.MultiDestination.WeightedDestination](#networking.zephyr.solo.io.TrafficPolicySpec.MultiDestination.WeightedDestination) | repeated |  |






<a name="networking.zephyr.solo.io.TrafficPolicySpec.MultiDestination.WeightedDestination"></a>

### TrafficPolicySpec.MultiDestination.WeightedDestination



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| destination | [core.zephyr.solo.io.ResourceRef](#core.zephyr.solo.io.ResourceRef) |  |  |
| weight | [uint32](#uint32) |  | Routing to each destination will be balanced by the ratio of the destination's weight to the total weight on a route. |
| subset | [][TrafficPolicySpec.MultiDestination.WeightedDestination.SubsetEntry](#networking.zephyr.solo.io.TrafficPolicySpec.MultiDestination.WeightedDestination.SubsetEntry) | repeated | Subset routing is currently only supported on Istio. |
| port | [uint32](#uint32) |  | Port on the destination service to receive traffic. If multiple are found, and none are specified, then the configuration will be considered invalid |






<a name="networking.zephyr.solo.io.TrafficPolicySpec.MultiDestination.WeightedDestination.SubsetEntry"></a>

### TrafficPolicySpec.MultiDestination.WeightedDestination.SubsetEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="networking.zephyr.solo.io.TrafficPolicySpec.QueryParameterMatcher"></a>

### TrafficPolicySpec.QueryParameterMatcher
Query parameter matching treats the query string of a request's :path header as an ampersand-separated list of keys and/or key=value elements.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Specifies the name of a key that must be present in the requested path*'s query string. |
| value | [string](#string) |  | Specifies the value of the key. If the value is absent, a request that contains the key in its query string will match, whether the key appears with a value (e.g., "?debug=true") or not (e.g., "?debug") |
| regex | [bool](#bool) |  | Specifies whether the query parameter value is a regular expression. Defaults to false. The entire query parameter value (i.e., the part to the right of the equals sign in "key=value") must match the regex. E.g., the regex "\d+$" will match "123" but not "a123" or "123a". |






<a name="networking.zephyr.solo.io.TrafficPolicySpec.RetryPolicy"></a>

### TrafficPolicySpec.RetryPolicy
RetryPolicy contains mesh-specific retry configuration Different meshes support different Retry features Service Mesh Hub's RetryPolicy exposes config for multiple meshes simultaneously, Allowing the same TrafficPolicy to apply retries to different mesh types The configuration applied to the target mesh will use the corresponding config for each type, while other config types will be ignored


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| attempts | [int32](#int32) |  | Number of retries for a given request |
| perTryTimeout | [google.protobuf.Duration](#google.protobuf.Duration) |  | Timeout per retry attempt for a given request. format: 1h/1m/1s/1ms. MUST BE >=1ms. |






<a name="networking.zephyr.solo.io.TrafficPolicySpec.StringMatch"></a>

### TrafficPolicySpec.StringMatch
Describes how to match a given string in HTTP headers. Match is case-sensitive.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exact | [string](#string) |  | Exact string match. |
| prefix | [string](#string) |  | Prefix-based match. |
| regex | [string](#string) |  | ECMAscript style regex-based match. |






<a name="networking.zephyr.solo.io.TrafficPolicyStatus"></a>

### TrafficPolicyStatus



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| translationStatus | [core.zephyr.solo.io.Status](#core.zephyr.solo.io.Status) |  |  |
| translatorErrors | [][TrafficPolicyStatus.TranslatorError](#networking.zephyr.solo.io.TrafficPolicyStatus.TranslatorError) | repeated |  |






<a name="networking.zephyr.solo.io.TrafficPolicyStatus.TranslatorError"></a>

### TrafficPolicyStatus.TranslatorError



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| translatorId | [string](#string) |  | ID representing a translator that translates TrafficPolicy to Mesh-specific config. |
| errorMessage | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->

