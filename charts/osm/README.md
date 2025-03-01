# Open Service Mesh Edge Helm Chart

![Version: 1.2.0](https://img.shields.io/badge/Version-1.2.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: v1.2.0](https://img.shields.io/badge/AppVersion-v1.2.0-informational?style=flat-square)

A Helm chart to install the [osm-edge](https://github.com/flomesh-io/osm-edge) control plane on Kubernetes.

## Prerequisites

- Kubernetes >= 1.19.0-0

## Get Repo Info

```console
helm repo add osm https://flomesh-io.github.io/osm-edge
helm repo update
```

## Install Chart

```console
helm install [RELEASE_NAME] osm/osm
```

The command deploys `osm-controller` on the Kubernetes cluster in the default configuration.

_See [configuration](#configuration) below._

_See [helm install](https://helm.sh/docs/helm/helm_install/) for command documentation._

## Uninstall Chart

```console
helm uninstall [RELEASE_NAME]
```

This removes all the Kubernetes components associated with the chart and deletes the release.

_See [helm uninstall](https://helm.sh/docs/helm/helm_uninstall/) for command documentation._

## Upgrading Chart

```console
helm upgrade [RELEASE_NAME] [CHART] --install
```

_See [helm upgrade](https://helm.sh/docs/helm/helm_upgrade/) for command documentation._

## Configuration

See [Customizing the Chart Before Installing](https://helm.sh/docs/intro/using_helm/#customizing-the-chart-before-installing). To see all configurable options with detailed comments, visit the chart's [values.yaml](./values.yaml), or run these configuration commands:

```console
helm show values osm/osm
```

The following table lists the configurable parameters of the osm chart and their default values.

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| contour.contour | object | `{"image":{"registry":"ghcr.io","repository":"projectcontour/contour","tag":"v1.21.1"}}` | Contour controller configuration |
| contour.enabled | bool | `false` | Enables deployment of Contour control plane and gateway |
| contour.envoy | object | `{"image":{"registry":"docker.io","repository":"envoyproxy/envoy-distroless","tag":"v1.22.2"}}` | Contour envoy edge proxy configuration |
| fsm.enabled | bool | `false` | Enables deployment of fsm control plane and gateway |
| osm.caBundleSecretName | string | `"osm-ca-bundle"` | The Kubernetes secret name to store CA bundle for the root CA used in OSM |
| osm.certificateProvider.certKeyBitSize | int | `2048` | Certificate key bit size for data plane certificates issued to workloads to communicate over mTLS |
| osm.certificateProvider.kind | string | `"tresor"` | The Certificate manager type: `tresor`, `vault` or `cert-manager` |
| osm.certificateProvider.serviceCertValidityDuration | string | `"24h"` | Service certificate validity duration for certificate issued to workloads to communicate over mTLS |
| osm.certmanager.issuerGroup | string | `"cert-manager.io"` | cert-manager issuer group |
| osm.certmanager.issuerKind | string | `"Issuer"` | cert-manager issuer kind |
| osm.certmanager.issuerName | string | `"osm-ca"` | cert-manager issuer namecert-manager issuer name |
| osm.cleanup.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[0].key | string | `"kubernetes.io/os"` |  |
| osm.cleanup.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[0].operator | string | `"In"` |  |
| osm.cleanup.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[0].values[0] | string | `"linux"` |  |
| osm.cleanup.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[1].key | string | `"kubernetes.io/arch"` |  |
| osm.cleanup.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[1].operator | string | `"In"` |  |
| osm.cleanup.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[1].values[0] | string | `"amd64"` |  |
| osm.cleanup.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[1].values[1] | string | `"arm64"` |  |
| osm.cleanup.nodeSelector | object | `{}` |  |
| osm.cleanup.tolerations | list | `[]` | Node tolerations applied to control plane pods. The specified tolerations allow pods to schedule onto nodes with matching taints. |
| osm.configResyncInterval | string | `"0s"` | Sets the resync interval for regular proxy broadcast updates, set to 0s to not enforce any resync |
| osm.controlPlaneTolerations | list | `[]` | Node tolerations applied to control plane pods. The specified tolerations allow pods to schedule onto nodes with matching taints. |
| osm.controllerLogLevel | string | `"info"` | Controller log verbosity |
| osm.curlImage | string | `"curlimages/curl"` | Curl image for control plane init container |
| osm.deployGrafana | bool | `false` | Deploy Grafana with OSM installation |
| osm.deployJaeger | bool | `false` | Deploy Jaeger during OSM installation |
| osm.deployPrometheus | bool | `false` | Deploy Prometheus with OSM installation |
| osm.enableDebugServer | bool | `false` | Enable the debug HTTP server on OSM controller |
| osm.enableEgress | bool | `true` | Enable egress in the mesh |
| osm.enableFluentbit | bool | `false` | Enable Fluent Bit sidecar deployment on OSM controller's pod |
| osm.enablePermissiveTrafficPolicy | bool | `true` | Enable permissive traffic policy mode |
| osm.enablePrivilegedInitContainer | bool | `false` | Run init container in privileged mode |
| osm.enableReconciler | bool | `false` | Enable reconciler for OSM's CRDs and mutating webhook |
| osm.enforceSingleMesh | bool | `true` | Enforce only deploying one mesh in the cluster |
| osm.featureFlags.enableAccessCertPolicy | bool | `false` |  |
| osm.featureFlags.enableAccessControlPolicy | bool | `false` | Enables OSM's AccessControl policy API. When enabled, OSM will use the AccessControl API allow access control traffic to mesh backends |
| osm.featureFlags.enableAsyncProxyServiceMapping | bool | `false` | Enable async proxy-service mapping |
| osm.featureFlags.enableEgressPolicy | bool | `true` | Enable OSM's Egress policy API. When enabled, fine grained control over Egress (external) traffic is enforced |
| osm.featureFlags.enableIngressBackendPolicy | bool | `true` | Enables OSM's IngressBackend policy API. When enabled, OSM will use the IngressBackend API allow ingress traffic to mesh backends |
| osm.featureFlags.enableMeshRootCertificate | bool | `false` | Enable the MeshRootCertificate to configure the OSM certificate provider |
| osm.featureFlags.enableRetryPolicy | bool | `false` | Enable Retry Policy for automatic request retries |
| osm.featureFlags.enableSidecarActiveHealthChecks | bool | `false` | Enable Sidecar active health checks |
| osm.featureFlags.enableSnapshotCacheMode | bool | `false` | Enables SnapshotCache feature for Sidecar xDS server. |
| osm.featureFlags.enableWASMStats | bool | `true` | Enable extra Envoy statistics generated by a custom WASM extension |
| osm.fluentBit.enableProxySupport | bool | `false` | Enable proxy support toggle for Fluent Bit |
| osm.fluentBit.httpProxy | string | `""` | Optional HTTP proxy endpoint for Fluent Bit |
| osm.fluentBit.httpsProxy | string | `""` | Optional HTTPS proxy endpoint for Fluent Bit |
| osm.fluentBit.name | string | `"fluentbit-logger"` | Fluent Bit sidecar container name |
| osm.fluentBit.outputPlugin | string | `"stdout"` | Fluent Bit output plugin |
| osm.fluentBit.primaryKey | string | `""` | Primary Key for Fluent Bit output plugin to Log Analytics |
| osm.fluentBit.pullPolicy | string | `"IfNotPresent"` | PullPolicy for Fluent Bit sidecar container |
| osm.fluentBit.registry | string | `"fluent"` | Registry for Fluent Bit sidecar container |
| osm.fluentBit.tag | string | `"1.6.4"` | Fluent Bit sidecar image tag |
| osm.fluentBit.workspaceId | string | `""` | WorkspaceId for Fluent Bit output plugin to Log Analytics |
| osm.grafana.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[0].key | string | `"kubernetes.io/os"` |  |
| osm.grafana.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[0].operator | string | `"In"` |  |
| osm.grafana.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[0].values[0] | string | `"linux"` |  |
| osm.grafana.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[1].key | string | `"kubernetes.io/arch"` |  |
| osm.grafana.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[1].operator | string | `"In"` |  |
| osm.grafana.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[1].values[0] | string | `"amd64"` |  |
| osm.grafana.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[1].values[1] | string | `"arm64"` |  |
| osm.grafana.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[1].values[2] | string | `"arm"` |  |
| osm.grafana.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[1].values[3] | string | `"ppc64le"` |  |
| osm.grafana.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[1].values[4] | string | `"s390x"` |  |
| osm.grafana.enableRemoteRendering | bool | `false` | Enable Remote Rendering in Grafana |
| osm.grafana.image | string | `"grafana/grafana:8.2.2"` | Image used for Grafana |
| osm.grafana.nodeSelector | object | `{}` |  |
| osm.grafana.port | int | `3000` | Grafana service's port |
| osm.grafana.rendererImage | string | `"grafana/grafana-image-renderer:3.2.1"` | Image used for Grafana Renderer |
| osm.grafana.tolerations | list | `[]` | Node tolerations applied to control plane pods. The specified tolerations allow pods to schedule onto nodes with matching taints. |
| osm.image.digest | object | `{"osmBootstrap":"","osmCRDs":"","osmController":"","osmHealthcheck":"","osmInjector":"","osmPreinstall":"","osmSidecarInit":""}` | Image digest (defaults to latest compatible tag) |
| osm.image.digest.osmBootstrap | string | `""` | osm-boostrap's image digest |
| osm.image.digest.osmCRDs | string | `""` | osm-crds' image digest |
| osm.image.digest.osmController | string | `""` | osm-controller's image digest |
| osm.image.digest.osmHealthcheck | string | `""` | osm-healthcheck's image digest |
| osm.image.digest.osmInjector | string | `""` | osm-injector's image digest |
| osm.image.digest.osmPreinstall | string | `""` | osm-preinstall's image digest |
| osm.image.digest.osmSidecarInit | string | `""` | Sidecar init container's image digest |
| osm.image.name | object | `{"osmBootstrap":"osm-edge-bootstrap","osmCRDs":"osm-edge-crds","osmController":"osm-edge-controller","osmHealthcheck":"osm-edge-healthcheck","osmInjector":"osm-edge-injector","osmPreinstall":"osm-edge-preinstall","osmSidecarInit":"osm-edge-sidecar-init"}` | Image name defaults |
| osm.image.name.osmBootstrap | string | `"osm-edge-bootstrap"` | osm-boostrap's image name |
| osm.image.name.osmCRDs | string | `"osm-edge-crds"` | osm-crds' image name |
| osm.image.name.osmController | string | `"osm-edge-controller"` | osm-controller's image name |
| osm.image.name.osmHealthcheck | string | `"osm-edge-healthcheck"` | osm-healthcheck's image name |
| osm.image.name.osmInjector | string | `"osm-edge-injector"` | osm-injector's image name |
| osm.image.name.osmPreinstall | string | `"osm-edge-preinstall"` | osm-preinstall's image name |
| osm.image.name.osmSidecarInit | string | `"osm-edge-sidecar-init"` | Sidecar init container's image name |
| osm.image.pullPolicy | string | `"IfNotPresent"` | Container image pull policy for control plane containers |
| osm.image.registry | string | `"flomesh"` | Container image registry for control plane images |
| osm.image.tag | string | `"1.2.1"` | Container image tag for control plane images |
| osm.imagePullSecrets | list | `[]` | `osm-controller` image pull secret |
| osm.inboundPortExclusionList | list | `[]` | Specifies a global list of ports to exclude from inbound traffic interception by the sidecar proxy. If specified, must be a list of positive integers. |
| osm.injector.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[0].key | string | `"kubernetes.io/os"` |  |
| osm.injector.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[0].operator | string | `"In"` |  |
| osm.injector.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[0].values[0] | string | `"linux"` |  |
| osm.injector.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[1].key | string | `"kubernetes.io/arch"` |  |
| osm.injector.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[1].operator | string | `"In"` |  |
| osm.injector.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[1].values[0] | string | `"amd64"` |  |
| osm.injector.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[1].values[1] | string | `"arm64"` |  |
| osm.injector.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution[0].podAffinityTerm.labelSelector.matchExpressions[0].key | string | `"app"` |  |
| osm.injector.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution[0].podAffinityTerm.labelSelector.matchExpressions[0].operator | string | `"In"` |  |
| osm.injector.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution[0].podAffinityTerm.labelSelector.matchExpressions[0].values[0] | string | `"osm-injector"` |  |
| osm.injector.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution[0].podAffinityTerm.topologyKey | string | `"kubernetes.io/hostname"` |  |
| osm.injector.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution[0].weight | int | `100` |  |
| osm.injector.autoScale | object | `{"cpu":{"targetAverageUtilization":80},"enable":false,"maxReplicas":5,"memory":{"targetAverageUtilization":80},"minReplicas":1}` | Auto scale configuration |
| osm.injector.autoScale.cpu.targetAverageUtilization | int | `80` | Average target CPU utilization (%) |
| osm.injector.autoScale.enable | bool | `false` | Enable Autoscale |
| osm.injector.autoScale.maxReplicas | int | `5` | Maximum replicas for autoscale |
| osm.injector.autoScale.memory.targetAverageUtilization | int | `80` | Average target memory utilization (%) |
| osm.injector.autoScale.minReplicas | int | `1` | Minimum replicas for autoscale |
| osm.injector.enablePodDisruptionBudget | bool | `false` | Enable Pod Disruption Budget |
| osm.injector.nodeSelector | object | `{}` |  |
| osm.injector.podLabels | object | `{}` | Sidecar injector's pod labels |
| osm.injector.replicaCount | int | `1` | Sidecar injector's replica count (ignored when autoscale.enable is true) |
| osm.injector.resource | object | `{"limits":{"cpu":"0.5","memory":"64M"},"requests":{"cpu":"0.3","memory":"64M"}}` | Sidecar injector's container resource parameters |
| osm.injector.tolerations | list | `[]` | Node tolerations applied to control plane pods. The specified tolerations allow pods to schedule onto nodes with matching taints. |
| osm.injector.webhookTimeoutSeconds | int | `20` | Mutating webhook timeout |
| osm.localProxyMode | string | `"Localhost"` | Proxy mode for the proxy sidecar. Acceptable values are ['Localhost', 'PodIP'] |
| osm.maxDataPlaneConnections | int | `0` | Sets the max data plane connections allowed for an instance of osm-controller, set to 0 to not enforce limits |
| osm.meshName | string | `"osm"` | Identifier for the instance of a service mesh within a cluster |
| osm.networkInterfaceExclusionList | list | `[]` | Specifies a global list of network interface names to exclude for inbound and outbound traffic interception by the sidecar proxy. |
| osm.osmBootstrap.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[0].key | string | `"kubernetes.io/os"` |  |
| osm.osmBootstrap.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[0].operator | string | `"In"` |  |
| osm.osmBootstrap.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[0].values[0] | string | `"linux"` |  |
| osm.osmBootstrap.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[1].key | string | `"kubernetes.io/arch"` |  |
| osm.osmBootstrap.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[1].operator | string | `"In"` |  |
| osm.osmBootstrap.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[1].values[0] | string | `"amd64"` |  |
| osm.osmBootstrap.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[1].values[1] | string | `"arm64"` |  |
| osm.osmBootstrap.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution[0].podAffinityTerm.labelSelector.matchExpressions[0].key | string | `"app"` |  |
| osm.osmBootstrap.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution[0].podAffinityTerm.labelSelector.matchExpressions[0].operator | string | `"In"` |  |
| osm.osmBootstrap.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution[0].podAffinityTerm.labelSelector.matchExpressions[0].values[0] | string | `"osm-bootstrap"` |  |
| osm.osmBootstrap.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution[0].podAffinityTerm.topologyKey | string | `"kubernetes.io/hostname"` |  |
| osm.osmBootstrap.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution[0].weight | int | `100` |  |
| osm.osmBootstrap.nodeSelector | object | `{}` |  |
| osm.osmBootstrap.podLabels | object | `{}` | OSM bootstrap's pod labels |
| osm.osmBootstrap.replicaCount | int | `1` | OSM bootstrap's replica count |
| osm.osmBootstrap.resource | object | `{"limits":{"cpu":"0.5","memory":"128M"},"requests":{"cpu":"0.3","memory":"128M"}}` | OSM bootstrap's container resource parameters |
| osm.osmBootstrap.tolerations | list | `[]` | Node tolerations applied to control plane pods. The specified tolerations allow pods to schedule onto nodes with matching taints. |
| osm.osmController.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[0].key | string | `"kubernetes.io/os"` |  |
| osm.osmController.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[0].operator | string | `"In"` |  |
| osm.osmController.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[0].values[0] | string | `"linux"` |  |
| osm.osmController.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[1].key | string | `"kubernetes.io/arch"` |  |
| osm.osmController.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[1].operator | string | `"In"` |  |
| osm.osmController.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[1].values[0] | string | `"amd64"` |  |
| osm.osmController.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[1].values[1] | string | `"arm64"` |  |
| osm.osmController.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution[0].podAffinityTerm.labelSelector.matchExpressions[0].key | string | `"app"` |  |
| osm.osmController.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution[0].podAffinityTerm.labelSelector.matchExpressions[0].operator | string | `"In"` |  |
| osm.osmController.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution[0].podAffinityTerm.labelSelector.matchExpressions[0].values[0] | string | `"osm-controller"` |  |
| osm.osmController.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution[0].podAffinityTerm.topologyKey | string | `"kubernetes.io/hostname"` |  |
| osm.osmController.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution[0].weight | int | `100` |  |
| osm.osmController.autoScale | object | `{"cpu":{"targetAverageUtilization":80},"enable":false,"maxReplicas":5,"memory":{"targetAverageUtilization":80},"minReplicas":1}` | Auto scale configuration |
| osm.osmController.autoScale.cpu.targetAverageUtilization | int | `80` | Average target CPU utilization (%) |
| osm.osmController.autoScale.enable | bool | `false` | Enable Autoscale |
| osm.osmController.autoScale.maxReplicas | int | `5` | Maximum replicas for autoscale |
| osm.osmController.autoScale.memory.targetAverageUtilization | int | `80` | Average target memory utilization (%) |
| osm.osmController.autoScale.minReplicas | int | `1` | Minimum replicas for autoscale |
| osm.osmController.enablePodDisruptionBudget | bool | `false` | Enable Pod Disruption Budget |
| osm.osmController.podLabels | object | `{}` | OSM controller's pod labels |
| osm.osmController.replicaCount | int | `1` | OSM controller's replica count (ignored when autoscale.enable is true) |
| osm.osmController.resource | object | `{"limits":{"cpu":"1.5","memory":"1G"},"requests":{"cpu":"0.5","memory":"128M"}}` | OSM controller's container resource parameters. See https://docs.openservicemesh.io/docs/guides/ha_scale/scale/ for more details. |
| osm.osmController.tolerations | list | `[]` | Node tolerations applied to control plane pods. The specified tolerations allow pods to schedule onto nodes with matching taints. |
| osm.osmNamespace | string | `""` | Namespace to deploy OSM in. If not specified, the Helm release namespace is used. |
| osm.outboundIPRangeExclusionList | list | `[]` | Specifies a global list of IP ranges to exclude from outbound traffic interception by the sidecar proxy. If specified, must be a list of IP ranges of the form a.b.c.d/x. |
| osm.outboundIPRangeInclusionList | list | `[]` | Specifies a global list of IP ranges to include for outbound traffic interception by the sidecar proxy. If specified, must be a list of IP ranges of the form a.b.c.d/x. |
| osm.outboundPortExclusionList | list | `[]` | Specifies a global list of ports to exclude from outbound traffic interception by the sidecar proxy. If specified, must be a list of positive integers. |
| osm.pipyRepoImage | string | `"flomesh/pipy-repo:0.70.0-46"` | Pipy repo image for Pipy sidecar's proxy control plane container |
| osm.preinstall.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[0].key | string | `"kubernetes.io/os"` |  |
| osm.preinstall.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[0].operator | string | `"In"` |  |
| osm.preinstall.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[0].values[0] | string | `"linux"` |  |
| osm.preinstall.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[1].key | string | `"kubernetes.io/arch"` |  |
| osm.preinstall.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[1].operator | string | `"In"` |  |
| osm.preinstall.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[1].values[0] | string | `"amd64"` |  |
| osm.preinstall.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[1].values[1] | string | `"arm64"` |  |
| osm.preinstall.nodeSelector | object | `{}` |  |
| osm.preinstall.tolerations | list | `[]` | Node tolerations applied to control plane pods. The specified tolerations allow pods to schedule onto nodes with matching taints. |
| osm.prometheus.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[0].key | string | `"kubernetes.io/os"` |  |
| osm.prometheus.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[0].operator | string | `"In"` |  |
| osm.prometheus.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[0].values[0] | string | `"linux"` |  |
| osm.prometheus.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[1].key | string | `"kubernetes.io/arch"` |  |
| osm.prometheus.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[1].operator | string | `"In"` |  |
| osm.prometheus.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[1].values[0] | string | `"amd64"` |  |
| osm.prometheus.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[1].values[1] | string | `"arm64"` |  |
| osm.prometheus.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[1].values[2] | string | `"arm"` |  |
| osm.prometheus.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[1].values[3] | string | `"ppc64le"` |  |
| osm.prometheus.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[1].values[4] | string | `"s390x"` |  |
| osm.prometheus.image | string | `"prom/prometheus:v2.34.0"` | Image used for Prometheus |
| osm.prometheus.nodeSelector | object | `{}` |  |
| osm.prometheus.port | int | `7070` | Prometheus service's port |
| osm.prometheus.resources | object | `{"limits":{"cpu":"1","memory":"2G"},"requests":{"cpu":"0.5","memory":"512M"}}` | Prometheus's container resource parameters |
| osm.prometheus.retention | object | `{"time":"15d"}` | Prometheus data rentention configuration |
| osm.prometheus.retention.time | string | `"15d"` | Prometheus data retention time |
| osm.prometheus.tolerations | list | `[]` | Node tolerations applied to control plane pods. The specified tolerations allow pods to schedule onto nodes with matching taints. |
| osm.remoteLogging.address | string | `""` | Address of the remote logging service (must contain the namespace). When left empty, this is computed in helper template to "remote-logging-service.<osm-namespace>.svc.cluster.local". |
| osm.remoteLogging.authorization | string | `""` | The authorization for remote logging service |
| osm.remoteLogging.enable | bool | `false` | Toggles Sidecar's remote logging functionality on/off for all sidecar proxies in the mesh |
| osm.remoteLogging.endpoint | string | `""` | Remote logging's API path where the spans will be sent to |
| osm.remoteLogging.port | int | `30514` | Port of the remote logging service |
| osm.sidecarClass | string | `"pipy"` | The class of the OSM Sidecar Driver |
| osm.sidecarDrivers | list | `[{"proxyServerPort":6060,"sidecarImage":"flomesh/pipy:0.70.0-46","sidecarName":"pipy"},{"proxyServerPort":15128,"sidecarImage":"envoyproxy/envoy:v1.19.3","sidecarName":"envoy","sidecarWindowsImage":"envoyproxy/envoy-windows:latest"}]` | Sidecar drivers supported by osm-edge |
| osm.sidecarDrivers[0].proxyServerPort | int | `6060` | Remote destination port on which the Discovery Service listens for new connections from Sidecars. |
| osm.sidecarDrivers[0].sidecarImage | string | `"flomesh/pipy:0.70.0-46"` | Sidecar image for Linux workloads |
| osm.sidecarDrivers[1].proxyServerPort | int | `15128` | Remote destination port on which the Discovery Service listens for new connections from Sidecars. |
| osm.sidecarDrivers[1].sidecarImage | string | `"envoyproxy/envoy:v1.19.3"` | Sidecar image for Linux workloads |
| osm.sidecarDrivers[1].sidecarWindowsImage | string | `"envoyproxy/envoy-windows:latest"` | Sidecar image for Windows workloads |
| osm.sidecarImage | string | `""` | Sidecar image for Linux workloads |
| osm.sidecarLogLevel | string | `"error"` | Log level for the proxy sidecar. Non developers should generally never set this value. In production environments the LogLevel should be set to `error` |
| osm.tracing.address | string | `""` | Address of the tracing collector service (must contain the namespace). When left empty, this is computed in helper template to "jaeger.<osm-namespace>.svc.cluster.local". Please override for BYO-tracing as documented in tracing.md |
| osm.tracing.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[0].key | string | `"kubernetes.io/os"` |  |
| osm.tracing.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[0].operator | string | `"In"` |  |
| osm.tracing.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[0].values[0] | string | `"linux"` |  |
| osm.tracing.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[1].key | string | `"kubernetes.io/arch"` |  |
| osm.tracing.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[1].operator | string | `"In"` |  |
| osm.tracing.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[1].values[0] | string | `"amd64"` |  |
| osm.tracing.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[1].values[1] | string | `"arm64"` |  |
| osm.tracing.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[1].values[2] | string | `"ppc64le"` |  |
| osm.tracing.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[0].matchExpressions[1].values[3] | string | `"s390x"` |  |
| osm.tracing.enable | bool | `false` | Toggles Sidecar's tracing functionality on/off for all sidecar proxies in the mesh |
| osm.tracing.endpoint | string | `"/api/v2/spans"` | Tracing collector's API path where the spans will be sent to |
| osm.tracing.image | string | `"jaegertracing/all-in-one"` | Image used for tracing |
| osm.tracing.nodeSelector | object | `{}` |  |
| osm.tracing.port | int | `9411` | Port of the tracing collector service |
| osm.tracing.tolerations | list | `[]` | Node tolerations applied to control plane pods. The specified tolerations allow pods to schedule onto nodes with matching taints. |
| osm.trustDomain | string | `"cluster.local"` | The trust domain to use as part of the common name when requesting new certificates. |
| osm.validatorWebhook.webhookConfigurationName | string | `""` | Name of the ValidatingWebhookConfiguration |
| osm.vault.host | string | `""` | Hashicorp Vault host/service - where Vault is installed |
| osm.vault.port | int | `8200` | port to use to connect to Vault |
| osm.vault.protocol | string | `"http"` | protocol to use to connect to Vault |
| osm.vault.role | string | `"openservicemesh"` | Vault role to be used by Open Service Mesh |
| osm.vault.secret | object | `{"key":"","name":""}` | The Kubernetes secret storing the Vault token used in OSM. The secret must be located in the namespace of the OSM installation |
| osm.vault.secret.key | string | `""` | The Kubernetes secret key with the value bring the Vault token |
| osm.vault.secret.name | string | `""` | The Kubernetes secret name storing the Vault token used in OSM |
| osm.vault.token | string | `""` | token that should be used to connect to Vault |
| osm.webhookConfigNamePrefix | string | `"osm-webhook"` | Prefix used in name of the webhook configuration resources |
| smi.validateTrafficTarget | bool | `true` | Enables validation of SMI Traffic Target |

<!-- markdownlint-enable MD013 MD034 -->
<!-- markdownlint-restore -->