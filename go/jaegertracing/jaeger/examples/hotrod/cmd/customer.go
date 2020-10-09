package cmd

import (
	"hotrod/pkg/log"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// customerCmd represents the customer command
var customerCmd = &cobra.Command{
	Use:   "customer",
	Short: "Starts Customer service",
	Long:  `Starts Customer service.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		zapLogger := logger.With(zap.String("service", "customer"))
		logger := log.NewFactory(zapLogger)
		server := customer.NewServer(
			
		)
	},
}
