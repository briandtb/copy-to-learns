package app

import "github.com/spf13/cobra"

// NewMetricsServerCommand provides a CLI handler for the metrics server entrypoint
func NewMetricsServerCommand(stopCh <-chan struct{}) *cobra.Command {
	opts := options.NewOptions()
}
