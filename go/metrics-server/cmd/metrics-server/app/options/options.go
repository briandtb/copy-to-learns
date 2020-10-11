package options

import (
	"fmt"
	"net"
	"time"

	"github.com/spf13/cobra"
	"k8s.io/apiserver/pkg/server"
	genericapiserver "k8s.io/apiserver/pkg/server"
	genericoptions "k8s.io/apiserver/pkg/server/options"
)

type Options struct {
	// genericoptions.ReccommendedOptions - EtcdOptions
	SecureServing  *genericoptions.SecureServingOptionsWithLoopback
	Authentication *genericoptions.DelegatingAuthenticationOptions
	Authorization  *genericoptions.DelegatingAuthorizationOptions
	Features       *genericoptions.FeatureOptions

	Kubeconfig string

	// Only to be used to for testing
	DisableAuthForTesting bool

	MetricResolution time.Duration

	KubeletUseNodeStatusPort     bool
	KubeletPort                  int
	InsecureKubeletTLS           bool
	KubeletPreferredAddressTypes []string
	KubeletCAFile                string
	KubeletClientKeyFile         string
	KubeletClientCertFile        string

	ShowVersion bool

	DeprecatedCompletelyInsecureKubelet bool
}

func (o *Options) Flags(cmd *cobra.Command) {
	flags := cmd.Flags()
	flags.DurationVar(&o.MetricResolution, "metric-resolution", o.MetricResolution, "The resolution at which metrics-server will retain metrics.")

	flags.BoolVar(&o.InsecureKubeletTLS, "kubelet-insecure-tls", o.InsecureKubeletTLS, "Do not verify CA of serving certificates presented by Kubelets. For testing purposes only.")
	flags.BoolVar(&o.DeprecatedCompletelyInsecureKubelet, "deprecated-kubelet-completely-insecure", o.DeprecatedCompletelyInsecureKubelet, "Do not use any encryption, authorization, or authentication when communiating with the Kubelet.")
	flags.BoolVar(&o.KubeletUseNodeStatusPort, "kubelet-use-node-status-port", o.KubeletUseNodeStatusPort, "Use the port in the node status. Takes precedence over --kubelet-port flags.")
	flags.IntVar(&o.KubeletPort, "kubelet-port", o.KubeletPort, "The port to use to connect to Kubelets.")
	flags.StringVar(&o.Kubeconfig, "kubeconfig", o.Kubeconfig, "The path to the kubeconfig used to connect to the Kubernetes API server and the Kubelets (default to in-cluster config)")
	flags.StringSliceVar(&o.KubeletPreferredAddressTypes, "kubelet-preferred-address-types", o.KubeletPreferredAddressTypes, "The priority of node address types to use when determining which address to use to connect to a particular node")
	flags.StringVar(&o.KubeletCAFile, "kubelet-certificate-authority", "", "Path to the CA to use to validate the Kubelet's serving certificates.")
	flags.StringVar(&o.KubeletClientKeyFile, "kubelet-client-key", "", "Path to a client key file for TLS.")
	flags.StringVar(&o.KubeletClientCertFile, "kubelet-client-certificate", "", "Path to a client cert file for TLS.")

	flags.BoolVar(&o.ShowVersion, "version", false, "Show version")

	flags.MarkDeprecated("deprecated-kubelet-completely-insecure", "This is rarely the right option, since it leaves kubelet communication completely insecure. If you encouter auth errors, make sure you've enabled token webhook auth on the Kubelet, and if you're in a test cluster with self-signed Kubelet certificates, consider using kubelet-insecure-tls instead.")

	o.SecureServing.AddFlags(flags)
	o.Authentication.AddFlags(flags)
	o.Authorization.AddFlags(flags)
	o.Features.AddFlags(flags)
}

// NewOptions constructs a new set of default options for metrics-server.
func NewOptions() *Options {
	o := &Options{
		SecureServing:  genericoptions.NewSecureServingOptions().WithLoopback(),
		Authentication: genericoptions.NewDelegatingAuthenticationOptions(),
		Authorization:  genericoptions.NewDelegatingAuthorizationOptions(),
		Features:       genericoptions.NewFeatureOptions(),

		MetricResolution:             60 * time.Second,
		KubeletPort:                  10250,
		KubeletPreferredAddressTypes: make([]string, len(utils.DefaultAddressTypePriority)),
	}

	for i, addrType := range utils.DefaultAddressTypePriority {
		o.KubeletPreferredAddressTypes[i] = string(addrType)
	}

	return o
}

func (o Options) ServerConfig() (*server.Config, error) {
	apiserver, err := o.ApiserverConfig()
}

func (o Options) ApiserverConfig() (*genericapiserver.Config, error) {
	if err := o.SecureServing.MaybeDefaultWithSelfSignedCerts("localhost", nil, []net.IP{net.ParseIP("127.0.0.1")}); err != nil {
		return nil, fmt.Errorf("error creating self-signed certificates: %v", err)
	}

	serverConfig := genericapiserver.NewConfig(api.Codecs)
	if err := o.SecureServing.ApplyTo(&serverConfig.SecureServing, &serverConfig.LoopbackClientConfig); err != nil {
		return nil, err
	}

	if !o.DisableAuthForTesting {
		if err := o.Authentication.ApplyTo(&serverConfig.Authentication, serverConfig.SecureServing, nil); err != nil {
			return nil, err
		}
		if err := o.Authorization.ApplyTo(&serverConfig.Authorization); err != nil {
			return nil, err
		}
	}
	serverConfig.Version = version.VersionInfo()
	// enable OpenAPI schemas
	serverConfig.OpenAPIConfig = genericapiserver.DefaultOpenAPIConfig(generatedopenapi.GetOpenAPIDefinitions, openapinamer.NewDefinitionNamer(api.Schema))

}
