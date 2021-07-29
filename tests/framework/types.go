package framework

import (
	"github.com/onsi/ginkgo"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/kind/pkg/cluster"

	"github.com/openservicemesh/osm/pkg/cli"
	versioned2 "github.com/openservicemesh/osm/pkg/gen/client/config/clientset/versioned"
	"github.com/openservicemesh/osm/pkg/gen/client/policy/clientset/versioned"
)

// OSMDescribeInfo is a struct to represent the Tier and Bucket of a given e2e test
type OSMDescribeInfo struct {
	// Tier represents the priority of the test. Lower value indicates higher priority.
	Tier int

	// Bucket indicates in which test Bucket the test will run in for CI. Each
	// Bucket is run in parallel while tests in the same Bucket run sequentially.
	Bucket int
}

// InstallType defines several OSM test deployment scenarios
type InstallType string

// CollectLogsType defines if/when to collect logs
type CollectLogsType string

// OsmTestData stores common state, variables and flags for the test at hand
type OsmTestData struct {
	T           ginkgo.GinkgoTInterface // for common test logging
	TestID      uint64                  // uint randomized for every test. GinkgoRandomSeed can't be used as is per-suite.
	TestDirBase string                  // Test directory base, default "/tmp", can be variable overridden
	TestDirName string                  // Autogenerated, based on test ID

	CleanupTest          bool            // Cleanup test-related resources once finished
	WaitForCleanup       bool            // Forces test to wait for effective deletion of resources upon cleanup
	IgnoreRestarts       bool            // Ignore control plane processes restarts, if any
	InstType             InstallType     // Install type
	CollectLogs          CollectLogsType // Collect logs type
	InitialRestartValues map[string]int  // Captures properly if an OSM instance have restarted during a NoInstall test

	OsmNamespace      string
	OsmMeshConfigName string
	OsmImageTag       string
	EnableNsMetricTag bool

	// Container registry related vars
	CtrRegistryUser     string // registry login
	CtrRegistryPassword string // registry password, if any
	CtrRegistryServer   string // server name. Has to be network reachable

	// Kind cluster related vars
	ClusterName                    string // Kind cluster name (used if kindCluster)
	CleanupKindClusterBetweenTests bool   // Clean and re-create kind cluster between tests
	CleanupKindCluster             bool   // Cleanup kind cluster upon test finish
	ClusterVersion                 string // Kind cluster version, ex. v1.20.2

	// Cluster handles and rest config
	Env        *cli.EnvSettings
	RestConfig *rest.Config
	Client     *kubernetes.Clientset

	SmiClients *smiClients

	// OSM's API clients
	PolicyClient *versioned.Clientset
	ConfigClient *versioned2.Clientset

	ClusterProvider *cluster.Provider // provider, used when kindCluster is used

	DeployOnOpenShift bool // Determines whether to configure tests for OpenShift
}

// InstallOSMOpts describes install options for OSM
type InstallOSMOpts struct {
	ControlPlaneNS          string
	CertManager             string
	ContainerRegistryLoc    string
	ContainerRegistrySecret string
	OsmImagetag             string
	DeployGrafana           bool
	DeployPrometheus        bool
	DeployJaeger            bool
	DeployFluentbit         bool

	VaultHost     string
	VaultProtocol string
	VaultToken    string
	VaultRole     string

	CertmanagerIssuerGroup string
	CertmanagerIssuerKind  string
	CertmanagerIssuerName  string
	CertKeyBitSize         int

	EgressEnabled        bool
	EnablePermissiveMode bool
	OSMLogLevel          string
	EnvoyLogLevel        string
	EnableDebugServer    bool

	SetOverrides []string

	EnablePrivilegedInitContainer bool
}

// CleanupType identifies what triggered the cleanup
type CleanupType string

//DockerConfig and other configs are docker-specific container registry secret structures.
// Most of it is taken or referenced from kubectl source itself
type DockerConfig map[string]DockerConfigEntry

// DockerConfigJSON  is a struct for docker-specific config
type DockerConfigJSON struct {
	Auths       DockerConfig      `json:"auths"`
	HTTPHeaders map[string]string `json:"HttpHeaders,omitempty"`
}

// DockerConfigEntry is a struct for docker-specific container registry secret structures
type DockerConfigEntry struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
	Auth     string `json:"auth,omitempty"`
}

// SuccessFunction is a simple definition for a success function.
// True as success, false otherwise
type SuccessFunction func() bool

// RetryOnErrorFunc is a function type passed to RetryFuncOnError() to execute
type RetryOnErrorFunc func() error
