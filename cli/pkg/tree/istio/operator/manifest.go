package operator

import (
	"bytes"
	"html/template"

	"github.com/solo-io/mesh-projects/cli/pkg/tree/istio/operator/install"
)

//go:generate mockgen -source ./manifest.go -destination mocks/mock_manifest_builder.go
type InstallerManifestBuilder interface {
	// Based on the pending installation config, generate an appropriate installation manifest
	Build(options *install.InstallationConfig) (installationManifest string, err error)

	// Generate an IstioControlPlane spec that sets up Istio with its demo profile
	GetControlPlaneSpecWithProfile(profile, installationNamespace string) (string, error)
}

func NewInstallerManifestBuilder() InstallerManifestBuilder {
	return &installerManifestBuilder{}
}

type installerManifestBuilder struct{}

func (i *installerManifestBuilder) Build(options *install.InstallationConfig) (string, error) {
	tmpl := template.New("")
	tmpl, err := tmpl.Parse(installationManifestTemplate)
	if err != nil {
		return "", err
	}

	var buffer bytes.Buffer

	err = tmpl.Execute(&buffer, options)
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}

func (i *installerManifestBuilder) GetControlPlaneSpecWithProfile(profile, namespace string) (string, error) {
	tmpl := template.New("")
	tmpl, err := tmpl.Parse(istioControlPlaneWithProfile)
	if err != nil {
		return "", err
	}

	var buffer bytes.Buffer

	err = tmpl.Execute(&buffer, &controlPlaneData{
		Profile:          profile,
		InstallNamespace: namespace,
	})
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}

type controlPlaneData struct {
	Profile          string
	InstallNamespace string
}

// the raw yaml was obtained from `https://istio.io/operator.yaml` as suggested by https://preliminary.istio.io/docs/setup/install/standalone-operator/
var installationManifestTemplate = `
{{- if .CreateNamespace }}
---
apiVersion: v1
kind: Namespace
metadata:
  name: {{ .InstallNamespace }}
...
{{- end }}
{{- if .CreateIstioControlPlaneCRD }}
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: istiocontrolplanes.install.istio.io
spec:
  group: install.istio.io
  names:
    kind: IstioControlPlane
    listKind: IstioControlPlaneList
    plural: istiocontrolplanes
    singular: istiocontrolplane
    shortNames:
    - icp
  scope: Namespaced
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation
            of an object. Servers should convert recognized schemas to the latest
            internal value, and may reject unrecognized values.
            More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#resources'
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this
            object represents. Servers may infer this from the endpoint the client
            submits requests to. Cannot be updated. In CamelCase.
            More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
          type: string
        spec:
          description: 'Specification of the desired state of the istio control plane resource.
            More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#spec-and-status'
          type: object
        status:
          description: 'Status describes each of istio control plane component status at the current time.
            0 means NONE, 1 means UPDATING, 2 means HEALTHY, 3 means ERROR, 4 means RECONCILING.
            More info: https://github.com/istio/operator/blob/master/pkg/apis/istio/v1alpha2/v1alpha2.pb.html &
            https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#spec-and-status'
          type: object
  versions:
  - name: v1alpha2
    served: true
    storage: true
...
{{- end }}
---
apiVersion: v1
kind: ServiceAccount
metadata:
  namespace: {{ .InstallNamespace }}
  name: istio-operator
...
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  creationTimestamp: null
  name: istio-operator
rules:
# istio groups
- apiGroups:
  - authentication.istio.io
  resources:
  - '*'
  verbs:
  - '*'
- apiGroups:
  - config.istio.io
  resources:
  - '*'
  verbs:
  - '*'
- apiGroups:
  - install.istio.io
  resources:
  - '*'
  verbs:
  - '*'
- apiGroups:
  - networking.istio.io
  resources:
  - '*'
  verbs:
  - '*'
- apiGroups:
  - rbac.istio.io
  resources:
  - '*'
  verbs:
  - '*'
- apiGroups:
  - security.istio.io
  resources:
  - '*'
  verbs:
  - '*'
# k8s groups
- apiGroups:
  - admissionregistration.k8s.io
  resources:
  - mutatingwebhookconfigurations
  - validatingwebhookconfigurations
  verbs:
  - '*'
- apiGroups:
  - apiextensions.k8s.io
  resources:
  - customresourcedefinitions.apiextensions.k8s.io
  - customresourcedefinitions
  verbs:
  - '*'
- apiGroups:
  - apps
  - extensions
  resources:
  - daemonsets
  - deployments
  - deployments/finalizers
  - ingresses
  - replicasets
  - statefulsets
  verbs:
  - '*'
- apiGroups:
  - autoscaling
  resources:
  - horizontalpodautoscalers
  verbs:
  - '*'
- apiGroups:
  - monitoring.coreos.com
  resources:
  - servicemonitors
  verbs:
  - get
  - create
- apiGroups:
  - policy
  resources:
  - poddisruptionbudgets
  verbs:
  - '*'
- apiGroups:
  - rbac.authorization.k8s.io
  resources:
  - clusterrolebindings
  - clusterroles
  - roles
  - rolebindings
  verbs:
  - '*'
- apiGroups:
  - ""
  resources:
  - configmaps  
  - endpoints
  - events
  - namespaces
  - pods
  - persistentvolumeclaims
  - secrets
  - services
  - serviceaccounts  
  verbs:
  - '*'
...
---
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: istio-operator
subjects:
- kind: ServiceAccount
  name: istio-operator
  namespace: {{ .InstallNamespace }}
roleRef:
  kind: ClusterRole
  name: istio-operator
  apiGroup: rbac.authorization.k8s.io
...
---
apiVersion: v1
kind: Service
metadata:
  namespace: {{ .InstallNamespace }}
  labels:
    name: istio-operator
  name: istio-operator-metrics
spec:
  ports:
  - name: http-metrics
    port: 8383
    targetPort: 8383
  selector:
    name: istio-operator
...
---
apiVersion: apps/v1
kind: Deployment
metadata:
  namespace: {{ .InstallNamespace }}
  name: istio-operator
spec:
  replicas: 1
  selector:
    matchLabels:
      name: istio-operator
  template:
    metadata:
      labels:
        name: istio-operator
    spec:
      serviceAccountName: istio-operator
      containers:
        - name: istio-operator
          image: docker.io/istio/operator:{{ .IstioOperatorVersion }}
          command:
          - istio-operator
          - server
          imagePullPolicy: Always
          resources:
            limits:
              cpu: 200m
              memory: 256Mi
            requests:
              cpu: 50m
              memory: 128Mi
          env:
            - name: WATCH_NAMESPACE
              value: ""
            - name: LEADER_ELECTION_NAMESPACE
              valueFrom:
                fieldRef:
                  fieldPath: metadata.namespace
            - name: POD_NAME
              valueFrom:
                fieldRef:
                  fieldPath: metadata.name
            - name: OPERATOR_NAME
              value: "istio-operator"
...

`

var istioControlPlaneWithProfile = `
apiVersion: install.istio.io/v1alpha2
kind: IstioControlPlane
metadata:
  namespace: {{ .InstallNamespace }}
  name: istiocontrolplane-{{ .Profile }}
spec:
  profile: {{ .Profile }}

`