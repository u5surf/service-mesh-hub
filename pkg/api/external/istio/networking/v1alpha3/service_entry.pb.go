// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/mesh-projects/api/external/istio/networking/v1alpha3/service_entry.proto

// `ServiceEntry` enables adding additional entries into Istio's internal_watcher
// service registry, so that auto-discovered services in the mesh can
// access/route to these manually specified services. A service entry
// describes the properties of a service (DNS name, VIPs, ports, protocols,
// endpoints). These services could be external to the mesh (e.g., web
// APIs) or mesh-internal_watcher services that are not part of the platform's
// service registry (e.g., a set of VMs talking to services in Kubernetes).
//
// The following example declares a few external APIs accessed by internal_watcher
// applications over HTTPS. The sidecar inspects the SNI value in the
// ClientHello message to route to the appropriate external service.
//
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: ServiceEntry
// metadata:
//   name: external-svc-https
// spec:
//   hosts:
//   - api.dropboxapi.com
//   - www.googleapis.com
//   - api.facebook.com
//   location: MESH_EXTERNAL
//   ports:
//   - number: 443
//     name: https
//     protocol: TLS
//   resolution: DNS
// ```
//
// The following configuration adds a set of MongoDB instances running on
// unmanaged VMs to Istio's registry, so that these services can be treated
// as any other service in the mesh. The associated DestinationRule is used
// to initiate mTLS connections to the database instances.
//
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: ServiceEntry
// metadata:
//   name: external-svc-mongocluster
// spec:
//   hosts:
//   - mymongodb.somedomain # not used
//   addresses:
//   - 192.192.192.192/24 # VIPs
//   ports:
//   - number: 27018
//     name: mongodb
//     protocol: MONGO
//   location: MESH_INTERNAL
//   resolution: STATIC
//   endpoints:
//   - address: 2.2.2.2
//   - address: 3.3.3.3
// ```
//
// and the associated DestinationRule
//
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: DestinationRule
// metadata:
//   name: mtls-mongocluster
// spec:
//   host: mymongodb.somedomain
//   trafficPolicy:
//     tls:
//       mode: MUTUAL
//       clientCertificate: /etc/certs/myclientcert.pem
//       privateKey: /etc/certs/client_private_key.pem
//       caCertificates: /etc/certs/rootcacerts.pem
// ```
//
// The following example uses a combination of service entry and TLS
// routing in a virtual service to steer traffic based on the SNI value to
// an internal_watcher egress firewall.
//
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: ServiceEntry
// metadata:
//   name: external-svc-redirect
// spec:
//   hosts:
//   - wikipedia.org
//   - "*.wikipedia.org"
//   location: MESH_EXTERNAL
//   ports:
//   - number: 443
//     name: https
//     protocol: TLS
//   resolution: NONE
// ```
//
// And the associated VirtualService to route based on the SNI value.
//
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: VirtualService
// metadata:
//   name: tls-routing
// spec:
//   hosts:
//   - wikipedia.org
//   - "*.wikipedia.org"
//   tls:
//   - match:
//     - sniHosts:
//       - wikipedia.org
//       - "*.wikipedia.org"
//     route:
//     - destination:
//         host: internal_watcher-egress-firewall.ns1.svc.cluster.local
// ```
//
// The virtual service with TLS match serves to override the default SNI
// match. In the absence of a virtual service, traffic will be forwarded to
// the wikipedia domains.
//
// The following example demonstrates the use of a dedicated egress gateway
// through which all external service traffic is forwarded.
// The 'exportTo' field allows for control over the visibility of a service
// declaration to other namespaces in the mesh. By default, a service is exported
// to all namespaces. The following example restricts the visibility to the
// current namespace, represented by ".", so that it cannot be used by other
// namespaces.
//
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: ServiceEntry
// metadata:
//   name: external-svc-httpbin
//   namespace : egress
// spec:
//   hosts:
//   - httpbin.com
//   exportTo:
//   - "."
//   location: MESH_EXTERNAL
//   ports:
//   - number: 80
//     name: http
//     protocol: HTTP
//   resolution: DNS
// ```
//
// Define a gateway to handle all egress traffic.
//
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: Gateway
// metadata:
//  name: istio-egressgateway
//  namespace: istio-system
// spec:
//  selector:
//    istio: egressgateway
//  servers:
//  - port:
//      number: 80
//      name: http
//      protocol: HTTP
//    hosts:
//    - "*"
// ```
//
// And the associated `VirtualService` to route from the sidecar to the
// gateway service (`istio-egressgateway.istio-system.svc.cluster.local`), as
// well as route from the gateway to the external service. Note that the
// virtual service is exported to all namespaces enabling them to route traffic
// through the gateway to the external service. Forcing traffic to go through
// a managed middle proxy like this is a common practice.
//
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: VirtualService
// metadata:
//   name: gateway-routing
//   namespace: egress
// spec:
//   hosts:
//   - httpbin.com
//   exportTo:
//   - "*"
//   gateways:
//   - mesh
//   - istio-egressgateway
//   http:
//   - match:
//     - port: 80
//       gateways:
//       - mesh
//     route:
//     - destination:
//         host: istio-egressgateway.istio-system.svc.cluster.local
//   - match:
//     - port: 80
//       gateways:
//       - istio-egressgateway
//     route:
//     - destination:
//         host: httpbin.com
// ```
//
// The following example demonstrates the use of wildcards in the hosts for
// external services. If the connection has to be routed to the IP address
// requested by the application (i.e. application resolves DNS and attempts
// to connect to a specific IP), the discovery mode must be set to `NONE`.
//
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: ServiceEntry
// metadata:
//   name: external-svc-wildcard-example
// spec:
//   hosts:
//   - "*.bar.com"
//   location: MESH_EXTERNAL
//   ports:
//   - number: 80
//     name: http
//     protocol: HTTP
//   resolution: NONE
// ```
//
// The following example demonstrates a service that is available via a
// Unix Domain Socket on the host of the client. The resolution must be
// set to STATIC to use Unix address endpoints.
//
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: ServiceEntry
// metadata:
//   name: unix-domain-socket-example
// spec:
//   hosts:
//   - "example.unix.local"
//   location: MESH_EXTERNAL
//   ports:
//   - number: 80
//     name: http
//     protocol: HTTP
//   resolution: STATIC
//   endpoints:
//   - address: unix:///var/run/example/socket
// ```
//
// For HTTP-based services, it is possible to create a `VirtualService`
// backed by multiple DNS addressable endpoints. In such a scenario, the
// application can use the `HTTP_PROXY` environment variable to transparently
// reroute API calls for the `VirtualService` to a chosen backend. For
// example, the following configuration creates a non-existent external
// service called foo.bar.com backed by three domains: us.foo.bar.com:8080,
// uk.foo.bar.com:9080, and in.foo.bar.com:7080
//
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: ServiceEntry
// metadata:
//   name: external-svc-dns
// spec:
//   hosts:
//   - foo.bar.com
//   location: MESH_EXTERNAL
//   ports:
//   - number: 80
//     name: http
//     protocol: HTTP
//   resolution: DNS
//   endpoints:
//   - address: us.foo.bar.com
//     ports:
//       https: 8080
//   - address: uk.foo.bar.com
//     ports:
//       https: 9080
//   - address: in.foo.bar.com
//     ports:
//       https: 7080
// ```
//
// With `HTTP_PROXY=http://localhost/`, calls from the application to
// `http://foo.bar.com` will be load balanced across the three domains
// specified above. In other words, a call to `http://foo.bar.com/baz` would
// be translated to `http://uk.foo.bar.com/baz`.
//
// The following example illustrates the usage of a `ServiceEntry`
// containing a subject alternate name
// whose format conforms to the [SPIFEE standard](https://github.com/spiffe/spiffe/blob/master/standards/SPIFFE-ID.md):
//
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: ServiceEntry
// metadata:
//   name: httpbin
//   namespace : httpbin-ns
// spec:
//   hosts:
//   - httpbin.com
//   location: MESH_INTERNAL
//   ports:
//   - number: 80
//     name: http
//     protocol: HTTP
//   resolution: STATIC
//   endpoints:
//   - address: 2.2.2.2
//   - address: 3.3.3.3
//   subjectAltNames:
//   - "spiffe://cluster.local/ns/httpbin-ns/sa/httpbin-service-account"
// ```
//

package v1alpha3

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Location specifies whether the service is part of Istio mesh or
// outside the mesh.  Location determines the behavior of several
// features, such as service-to-service mTLS authentication, policy
// enforcement, etc. When communicating with services outside the mesh,
// Istio's mTLS authentication is disabled, and policy enforcement is
// performed on the client-side as opposed to server-side.
type ServiceEntry_Location int32

const (
	// Signifies that the service is external to the mesh. Typically used
	// to indicate external services consumed through APIs.
	ServiceEntry_MESH_EXTERNAL ServiceEntry_Location = 0
	// Signifies that the service is part of the mesh. Typically used to
	// indicate services added explicitly as part of expanding the service
	// mesh to include unmanaged infrastructure (e.g., VMs added to a
	// Kubernetes based service mesh).
	ServiceEntry_MESH_INTERNAL ServiceEntry_Location = 1
)

var ServiceEntry_Location_name = map[int32]string{
	0: "MESH_EXTERNAL",
	1: "MESH_INTERNAL",
}

var ServiceEntry_Location_value = map[string]int32{
	"MESH_EXTERNAL": 0,
	"MESH_INTERNAL": 1,
}

func (x ServiceEntry_Location) String() string {
	return proto.EnumName(ServiceEntry_Location_name, int32(x))
}

func (ServiceEntry_Location) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e5504694f6dde3d4, []int{0, 0}
}

// Resolution determines how the proxy will resolve the IP addresses of
// the network endpoints associated with the service, so that it can
// route to one of them. The resolution mode specified here has no impact
// on how the application resolves the IP address associated with the
// service. The application may still have to use DNS to resolve the
// service to an IP so that the outbound traffic can be captured by the
// Proxy. Alternatively, for HTTP services, the application could
// directly communicate with the proxy (e.g., by setting HTTP_PROXY) to
// talk to these services.
type ServiceEntry_Resolution int32

const (
	// Assume that incoming connections have already been resolved (to a
	// specific destination IP address). Such connections are typically
	// routed via the proxy using mechanisms such as IP table REDIRECT/
	// eBPF. After performing any routing related transformations, the
	// proxy will forward the connection to the IP address to which the
	// connection was bound.
	ServiceEntry_NONE ServiceEntry_Resolution = 0
	// Use the static IP addresses specified in endpoints (see below) as the
	// backing instances associated with the service.
	ServiceEntry_STATIC ServiceEntry_Resolution = 1
	// Attempt to resolve the IP address by querying the ambient DNS,
	// during request processing. If no endpoints are specified, the proxy
	// will resolve the DNS address specified in the hosts field, if
	// wildcards are not used. If endpoints are specified, the DNS
	// addresses specified in the endpoints will be resolved to determine
	// the destination IP address.  DNS resolution cannot be used with Unix
	// domain socket endpoints.
	ServiceEntry_DNS ServiceEntry_Resolution = 2
)

var ServiceEntry_Resolution_name = map[int32]string{
	0: "NONE",
	1: "STATIC",
	2: "DNS",
}

var ServiceEntry_Resolution_value = map[string]int32{
	"NONE":   0,
	"STATIC": 1,
	"DNS":    2,
}

func (x ServiceEntry_Resolution) String() string {
	return proto.EnumName(ServiceEntry_Resolution_name, int32(x))
}

func (ServiceEntry_Resolution) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e5504694f6dde3d4, []int{0, 1}
}

type ServiceEntry struct {
	// Status indicates the validation status of this resource.
	// Status is read-only by clients, and set by supergloo during validation
	Status core.Status `protobuf:"bytes,100,opt,name=status,proto3" json:"status" testdiff:"ignore"`
	// Metadata contains the object metadata for this resource
	Metadata core.Metadata `protobuf:"bytes,101,opt,name=metadata,proto3" json:"metadata"`
	// REQUIRED. The hosts associated with the ServiceEntry. Could be a DNS
	// name with wildcard prefix.
	//
	// 1. The hosts field is used to select matching hosts in VirtualServices and DestinationRules.
	// 2. For HTTP traffic the HTTP Host/Authority header will be matched against the hosts field.
	// 3. For HTTPs or TLS traffic containing Server Name Indication (SNI), the SNI value
	// will be matched against the hosts field.
	//
	// Note that when resolution is set to type DNS
	// and no endpoints are specified, the host field will be used as the DNS name
	// of the endpoint to route traffic to.
	Hosts []string `protobuf:"bytes,1,rep,name=hosts,proto3" json:"hosts,omitempty"`
	// The virtual IP addresses associated with the service. Could be CIDR
	// prefix. For HTTP traffic, generated route configurations will include http route
	// domains for both the `addresses` and `hosts` field values and the destination will
	// be identified based on the HTTP Host/Authority header.
	// If one or more IP addresses are specified,
	// the incoming traffic will be identified as belonging to this service
	// if the destination IP matches the IP/CIDRs specified in the addresses
	// field. If the Addresses field is empty, traffic will be identified
	// solely based on the destination port. In such scenarios, the port on
	// which the service is being accessed must not be shared by any other
	// service in the mesh. In other words, the sidecar will behave as a
	// simple TCP proxy, forwarding incoming traffic on a specified port to
	// the specified destination endpoint IP/host. Unix domain socket
	// addresses are not supported in this field.
	Addresses []string `protobuf:"bytes,2,rep,name=addresses,proto3" json:"addresses,omitempty"`
	// REQUIRED. The ports associated with the external service. If the
	// Endpoints are Unix domain socket addresses, there must be exactly one
	// port.
	Ports []*Port `protobuf:"bytes,3,rep,name=ports,proto3" json:"ports,omitempty"`
	// Specify whether the service should be considered external to the mesh
	// or part of the mesh.
	Location ServiceEntry_Location `protobuf:"varint,4,opt,name=location,proto3,enum=istio.networking.v1alpha3.ServiceEntry_Location" json:"location,omitempty"`
	// REQUIRED: Service discovery mode for the hosts. Care must be taken
	// when setting the resolution mode to NONE for a TCP port without
	// accompanying IP addresses. In such cases, traffic to any IP on
	// said port will be allowed (i.e. 0.0.0.0:<port>).
	Resolution ServiceEntry_Resolution `protobuf:"varint,5,opt,name=resolution,proto3,enum=istio.networking.v1alpha3.ServiceEntry_Resolution" json:"resolution,omitempty"`
	// One or more endpoints associated with the service.
	Endpoints []*ServiceEntry_Endpoint `protobuf:"bytes,6,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	// A list of namespaces to which this service is exported. Exporting a service
	// allows it to be used by sidecars, gateways and virtual services defined in
	// other namespaces. This feature provides a mechanism for service owners
	// and mesh administrators to control the visibility of services across
	// namespace boundaries.
	//
	// If no namespaces are specified then the service is exported to all
	// namespaces by default.
	//
	// The value "." is reserved and defines an export to the same namespace that
	// the service is declared in. Similarly the value "*" is reserved and
	// defines an export to all namespaces.
	//
	// For a Kubernetes Service, the equivalent effect can be achieved by setting
	// the annotation "networking.istio.io/exportTo" to a comma-separated list
	// of namespace names.
	//
	// NOTE: in the current release, the `exportTo` value is restricted to
	// "." or "*" (i.e., the current namespace or all namespaces).
	ExportTo []string `protobuf:"bytes,7,rep,name=export_to,json=exportTo,proto3" json:"export_to,omitempty"`
	// The list of subject alternate names allowed for workload instances that
	// implement this service. This information is used to enforce
	// [secure-naming](https://istio.io/docs/concepts/security/#secure-naming).
	// If specified, the proxy will verify that the server
	// certificate's subject alternate name matches one of the specified values.
	SubjectAltNames      []string `protobuf:"bytes,8,rep,name=subject_alt_names,json=subjectAltNames,proto3" json:"subject_alt_names,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceEntry) Reset()         { *m = ServiceEntry{} }
func (m *ServiceEntry) String() string { return proto.CompactTextString(m) }
func (*ServiceEntry) ProtoMessage()    {}
func (*ServiceEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5504694f6dde3d4, []int{0}
}
func (m *ServiceEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceEntry.Unmarshal(m, b)
}
func (m *ServiceEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceEntry.Marshal(b, m, deterministic)
}
func (m *ServiceEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceEntry.Merge(m, src)
}
func (m *ServiceEntry) XXX_Size() int {
	return xxx_messageInfo_ServiceEntry.Size(m)
}
func (m *ServiceEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceEntry.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceEntry proto.InternalMessageInfo

func (m *ServiceEntry) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *ServiceEntry) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func (m *ServiceEntry) GetHosts() []string {
	if m != nil {
		return m.Hosts
	}
	return nil
}

func (m *ServiceEntry) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *ServiceEntry) GetPorts() []*Port {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *ServiceEntry) GetLocation() ServiceEntry_Location {
	if m != nil {
		return m.Location
	}
	return ServiceEntry_MESH_EXTERNAL
}

func (m *ServiceEntry) GetResolution() ServiceEntry_Resolution {
	if m != nil {
		return m.Resolution
	}
	return ServiceEntry_NONE
}

func (m *ServiceEntry) GetEndpoints() []*ServiceEntry_Endpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func (m *ServiceEntry) GetExportTo() []string {
	if m != nil {
		return m.ExportTo
	}
	return nil
}

func (m *ServiceEntry) GetSubjectAltNames() []string {
	if m != nil {
		return m.SubjectAltNames
	}
	return nil
}

// Endpoint defines a network address (IP or hostname) associated with
// the mesh service.
type ServiceEntry_Endpoint struct {
	// REQUIRED: Address associated with the network endpoint without the
	// port.  Domain names can be used if and only if the resolution is set
	// to DNS, and must be fully-qualified without wildcards. Use the form
	// unix:///absolute/path/to/socket for Unix domain socket endpoints.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// Set of ports associated with the endpoint. The ports must be
	// associated with a port name that was declared as part of the
	// service. Do not use for `unix://` addresses.
	Ports map[string]uint32 `protobuf:"bytes,2,rep,name=ports,proto3" json:"ports,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// One or more labels associated with the endpoint.
	Labels map[string]string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Network enables Istio to group endpoints resident in the same L3
	// domain/network. All endpoints in the same network are assumed to be
	// directly reachable from one another. When endpoints in different
	// networks cannot reach each other directly, an Istio Gateway can be
	// used to establish connectivity (usually using the
	// AUTO_PASSTHROUGH mode in a Gateway Server). This is
	// an advanced configuration used typically for spanning an Istio mesh
	// over multiple clusters.
	Network string `protobuf:"bytes,4,opt,name=network,proto3" json:"network,omitempty"`
	// The locality associated with the endpoint. A locality corresponds
	// to a failure domain (e.g., country/region/zone). Arbitrary failure
	// domain hierarchies can be represented by separating each
	// encapsulating failure domain by /. For example, the locality of an
	// an endpoint in US, in US-East-1 region, within availability zone
	// az-1, in data center rack r11 can be represented as
	// us/us-east-1/az-1/r11. Istio will configure the sidecar to route to
	// endpoints within the same locality as the sidecar. If none of the
	// endpoints in the locality are available, endpoints parent locality
	// (but within the same network ID) will be chosen. For example, if
	// there are two endpoints in same network (networkID "n1"), say e1
	// with locality us/us-east-1/az-1/r11 and e2 with locality
	// us/us-east-1/az-2/r12, a sidecar from us/us-east-1/az-1/r11 locality
	// will prefer e1 from the same locality over e2 from a different
	// locality. Endpoint e2 could be the IP associated with a gateway
	// (that bridges networks n1 and n2), or the IP associated with a
	// standard service endpoint.
	Locality string `protobuf:"bytes,5,opt,name=locality,proto3" json:"locality,omitempty"`
	// The load balancing weight associated with the endpoint. Endpoints
	// with higher weights will receive proportionally higher traffic.
	Weight               uint32   `protobuf:"varint,6,opt,name=weight,proto3" json:"weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceEntry_Endpoint) Reset()         { *m = ServiceEntry_Endpoint{} }
func (m *ServiceEntry_Endpoint) String() string { return proto.CompactTextString(m) }
func (*ServiceEntry_Endpoint) ProtoMessage()    {}
func (*ServiceEntry_Endpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5504694f6dde3d4, []int{0, 0}
}
func (m *ServiceEntry_Endpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceEntry_Endpoint.Unmarshal(m, b)
}
func (m *ServiceEntry_Endpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceEntry_Endpoint.Marshal(b, m, deterministic)
}
func (m *ServiceEntry_Endpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceEntry_Endpoint.Merge(m, src)
}
func (m *ServiceEntry_Endpoint) XXX_Size() int {
	return xxx_messageInfo_ServiceEntry_Endpoint.Size(m)
}
func (m *ServiceEntry_Endpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceEntry_Endpoint.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceEntry_Endpoint proto.InternalMessageInfo

func (m *ServiceEntry_Endpoint) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ServiceEntry_Endpoint) GetPorts() map[string]uint32 {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *ServiceEntry_Endpoint) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *ServiceEntry_Endpoint) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *ServiceEntry_Endpoint) GetLocality() string {
	if m != nil {
		return m.Locality
	}
	return ""
}

func (m *ServiceEntry_Endpoint) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func init() {
	proto.RegisterEnum("istio.networking.v1alpha3.ServiceEntry_Location", ServiceEntry_Location_name, ServiceEntry_Location_value)
	proto.RegisterEnum("istio.networking.v1alpha3.ServiceEntry_Resolution", ServiceEntry_Resolution_name, ServiceEntry_Resolution_value)
	proto.RegisterType((*ServiceEntry)(nil), "istio.networking.v1alpha3.ServiceEntry")
	proto.RegisterType((*ServiceEntry_Endpoint)(nil), "istio.networking.v1alpha3.ServiceEntry.Endpoint")
	proto.RegisterMapType((map[string]string)(nil), "istio.networking.v1alpha3.ServiceEntry.Endpoint.LabelsEntry")
	proto.RegisterMapType((map[string]uint32)(nil), "istio.networking.v1alpha3.ServiceEntry.Endpoint.PortsEntry")
}

func init() {
	proto.RegisterFile("github.com/solo-io/mesh-projects/api/external/istio/networking/v1alpha3/service_entry.proto", fileDescriptor_e5504694f6dde3d4)
}

var fileDescriptor_e5504694f6dde3d4 = []byte{
	// 707 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xd1, 0x4e, 0x13, 0x4d,
	0x14, 0x66, 0xdb, 0x52, 0x76, 0x0f, 0x3f, 0xfc, 0x65, 0x42, 0xc8, 0xd2, 0x9f, 0x40, 0xff, 0x5e,
	0x35, 0x1a, 0x76, 0xa1, 0xc4, 0x04, 0xd1, 0x1b, 0xd0, 0x46, 0x89, 0xa5, 0xea, 0xb4, 0x26, 0x46,
	0x2f, 0x9a, 0x69, 0x3b, 0x6c, 0xc7, 0x6e, 0x77, 0x9a, 0x9d, 0x69, 0x29, 0xb7, 0x3c, 0x81, 0x8f,
	0xe1, 0x23, 0xf8, 0x06, 0xfa, 0x14, 0x5c, 0xf8, 0x06, 0x98, 0x78, 0x6f, 0x76, 0x76, 0xb6, 0xad,
	0x46, 0x04, 0xee, 0xf6, 0x7c, 0xe7, 0x7c, 0x5f, 0xbf, 0x33, 0xe7, 0xf4, 0xc0, 0x7b, 0x8f, 0xc9,
	0xee, 0xb0, 0xe5, 0xb4, 0x79, 0xdf, 0x15, 0xdc, 0xe7, 0xdb, 0x8c, 0xbb, 0x7d, 0x2a, 0xba, 0xdb,
	0x83, 0x90, 0x7f, 0xa0, 0x6d, 0x29, 0x5c, 0x32, 0x60, 0x2e, 0x1d, 0x4b, 0x1a, 0x06, 0xc4, 0x77,
	0x99, 0x90, 0x8c, 0xbb, 0x01, 0x95, 0x67, 0x3c, 0xec, 0xb1, 0xc0, 0x73, 0x47, 0xbb, 0xc4, 0x1f,
	0x74, 0xc9, 0x9e, 0x2b, 0x68, 0x38, 0x62, 0x6d, 0xda, 0xa4, 0x81, 0x0c, 0xcf, 0x9d, 0x41, 0xc8,
	0x25, 0x47, 0xeb, 0xaa, 0xdc, 0x99, 0x96, 0x3b, 0x49, 0x79, 0xfe, 0xff, 0x3f, 0x69, 0x78, 0x44,
	0xd2, 0x33, 0xa2, 0xd9, 0xf9, 0x55, 0x8f, 0x7b, 0x5c, 0x7d, 0xba, 0xd1, 0x97, 0x46, 0x11, 0x1d,
	0xcb, 0x18, 0xa4, 0x63, 0xa9, 0xb1, 0x4d, 0xe5, 0xbc, 0xc7, 0xa4, 0x32, 0x3b, 0xda, 0x75, 0xfb,
	0x54, 0x92, 0x0e, 0x91, 0x44, 0xe7, 0x37, 0x7e, 0xcf, 0x0b, 0x49, 0xe4, 0x50, 0x5c, 0xc7, 0x4e,
	0xe2, 0x38, 0x5f, 0xfc, 0x62, 0xc2, 0x3f, 0xf5, 0xb8, 0xbb, 0x4a, 0xd4, 0x1c, 0x7a, 0x06, 0xd9,
	0x58, 0xc0, 0xee, 0x14, 0x8c, 0xd2, 0x62, 0x79, 0xd5, 0x69, 0xf3, 0x90, 0x3a, 0x11, 0xcd, 0x61,
	0xdc, 0xa9, 0xab, 0xdc, 0xd1, 0xfa, 0xd7, 0xcb, 0xad, 0xb9, 0xef, 0x97, 0x5b, 0x2b, 0x92, 0x0a,
	0xd9, 0x61, 0xa7, 0xa7, 0x07, 0x45, 0xe6, 0x05, 0x3c, 0xa4, 0x45, 0xac, 0xe9, 0x68, 0x1f, 0xcc,
	0xc4, 0xa9, 0x4d, 0x95, 0xd4, 0xda, 0xaf, 0x52, 0x27, 0x3a, 0x7b, 0x94, 0x89, 0xc4, 0xf0, 0xa4,
	0x1a, 0xad, 0xc2, 0x7c, 0x97, 0x0b, 0x29, 0x6c, 0xa3, 0x90, 0x2e, 0x59, 0x38, 0x0e, 0xd0, 0x06,
	0x58, 0xa4, 0xd3, 0x09, 0xa9, 0x10, 0x54, 0xd8, 0x29, 0x95, 0x99, 0x02, 0xe8, 0x01, 0xcc, 0x0f,
	0x78, 0x28, 0x85, 0x9d, 0x2e, 0xa4, 0x4b, 0x8b, 0xe5, 0x2d, 0xe7, 0xda, 0xe9, 0x38, 0xaf, 0x78,
	0x28, 0x71, 0x5c, 0x8d, 0xaa, 0x60, 0xfa, 0xbc, 0x4d, 0x24, 0xe3, 0x81, 0x9d, 0x29, 0x18, 0xa5,
	0xe5, 0xf2, 0xce, 0x5f, 0x98, 0xb3, 0x0f, 0xe5, 0x54, 0x35, 0x0f, 0x4f, 0x14, 0x10, 0x06, 0x08,
	0xa9, 0xe0, 0xfe, 0x50, 0xe9, 0xcd, 0x2b, 0xbd, 0xf2, 0x6d, 0xf5, 0xf0, 0x84, 0x89, 0x67, 0x54,
	0x50, 0x0d, 0x2c, 0x1a, 0x74, 0x06, 0x9c, 0x05, 0x52, 0xd8, 0x59, 0xd5, 0xdc, 0xad, 0x2d, 0x56,
	0x34, 0x11, 0x4f, 0x25, 0xd0, 0x7f, 0x60, 0xd1, 0x71, 0xd4, 0x7c, 0x53, 0x72, 0x7b, 0x41, 0x3d,
	0xa3, 0x19, 0x03, 0x0d, 0x8e, 0xee, 0xc1, 0x8a, 0x18, 0xb6, 0xa2, 0xbf, 0x46, 0x93, 0xf8, 0xb2,
	0x19, 0x90, 0x3e, 0x15, 0xb6, 0xa9, 0x8a, 0xfe, 0xd5, 0x89, 0x43, 0x5f, 0xd6, 0x22, 0x38, 0xff,
	0x31, 0x0d, 0x66, 0xf2, 0x03, 0xc8, 0x86, 0x05, 0x3d, 0x0b, 0xdb, 0x28, 0x18, 0x25, 0x0b, 0x27,
	0x21, 0x7a, 0x9d, 0x0c, 0x26, 0xa5, 0xbc, 0x3f, 0xba, 0xab, 0x77, 0x35, 0x2e, 0xa1, 0xb0, 0x64,
	0x68, 0x0d, 0xc8, 0xfa, 0xa4, 0x45, 0xfd, 0x64, 0xd8, 0x8f, 0xef, 0xac, 0x59, 0x55, 0xf4, 0x58,
	0x54, 0x6b, 0x45, 0x2d, 0x68, 0x01, 0xb5, 0x09, 0x16, 0x4e, 0x42, 0x94, 0x8f, 0x97, 0xc4, 0x67,
	0xf2, 0x5c, 0x0d, 0xd5, 0xc2, 0x93, 0x18, 0xad, 0x41, 0xf6, 0x8c, 0x32, 0xaf, 0x2b, 0xed, 0x6c,
	0xc1, 0x28, 0x2d, 0x61, 0x1d, 0xe5, 0xf7, 0x01, 0xa6, 0xc6, 0x51, 0x0e, 0xd2, 0x3d, 0x7a, 0xae,
	0x9f, 0x26, 0xfa, 0x8c, 0x76, 0x7c, 0x44, 0xfc, 0x21, 0xb5, 0x53, 0x8a, 0x16, 0x07, 0x07, 0xa9,
	0x7d, 0x23, 0xff, 0x10, 0x16, 0x67, 0xec, 0xdd, 0x44, 0xb5, 0x66, 0xa8, 0xc5, 0x1d, 0x30, 0x93,
	0xad, 0x44, 0x2b, 0xb0, 0x74, 0x52, 0xa9, 0x3f, 0x6f, 0x56, 0xde, 0x36, 0x2a, 0xb8, 0x76, 0x58,
	0xcd, 0xcd, 0x4d, 0xa0, 0xe3, 0x9a, 0x86, 0x8c, 0xe2, 0x7d, 0x80, 0xe9, 0xde, 0x21, 0x13, 0x32,
	0xb5, 0x97, 0xb5, 0x4a, 0x6e, 0x0e, 0x01, 0x64, 0xeb, 0x8d, 0xc3, 0xc6, 0xf1, 0x93, 0x9c, 0x81,
	0x16, 0x20, 0xfd, 0xb4, 0x56, 0xcf, 0xa5, 0x0e, 0x36, 0x2e, 0xae, 0x32, 0x19, 0x48, 0x09, 0x7a,
	0x71, 0x95, 0xc9, 0xa1, 0x65, 0x7d, 0x14, 0xa3, 0x9b, 0xc8, 0xa8, 0x38, 0x7a, 0xf3, 0xf9, 0x47,
	0xc6, 0xf8, 0xf4, 0x6d, 0xd3, 0x78, 0xf7, 0xe2, 0xc6, 0xb3, 0x3b, 0xe8, 0x79, 0xb7, 0x3c, 0xbd,
	0xad, 0xac, 0xba, 0x53, 0x7b, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xe6, 0xfd, 0xa1, 0xe9, 0xcc,
	0x05, 0x00, 0x00,
}

func (this *ServiceEntry) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ServiceEntry)
	if !ok {
		that2, ok := that.(ServiceEntry)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if len(this.Hosts) != len(that1.Hosts) {
		return false
	}
	for i := range this.Hosts {
		if this.Hosts[i] != that1.Hosts[i] {
			return false
		}
	}
	if len(this.Addresses) != len(that1.Addresses) {
		return false
	}
	for i := range this.Addresses {
		if this.Addresses[i] != that1.Addresses[i] {
			return false
		}
	}
	if len(this.Ports) != len(that1.Ports) {
		return false
	}
	for i := range this.Ports {
		if !this.Ports[i].Equal(that1.Ports[i]) {
			return false
		}
	}
	if this.Location != that1.Location {
		return false
	}
	if this.Resolution != that1.Resolution {
		return false
	}
	if len(this.Endpoints) != len(that1.Endpoints) {
		return false
	}
	for i := range this.Endpoints {
		if !this.Endpoints[i].Equal(that1.Endpoints[i]) {
			return false
		}
	}
	if len(this.ExportTo) != len(that1.ExportTo) {
		return false
	}
	for i := range this.ExportTo {
		if this.ExportTo[i] != that1.ExportTo[i] {
			return false
		}
	}
	if len(this.SubjectAltNames) != len(that1.SubjectAltNames) {
		return false
	}
	for i := range this.SubjectAltNames {
		if this.SubjectAltNames[i] != that1.SubjectAltNames[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ServiceEntry_Endpoint) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ServiceEntry_Endpoint)
	if !ok {
		that2, ok := that.(ServiceEntry_Endpoint)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Address != that1.Address {
		return false
	}
	if len(this.Ports) != len(that1.Ports) {
		return false
	}
	for i := range this.Ports {
		if this.Ports[i] != that1.Ports[i] {
			return false
		}
	}
	if len(this.Labels) != len(that1.Labels) {
		return false
	}
	for i := range this.Labels {
		if this.Labels[i] != that1.Labels[i] {
			return false
		}
	}
	if this.Network != that1.Network {
		return false
	}
	if this.Locality != that1.Locality {
		return false
	}
	if this.Weight != that1.Weight {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
