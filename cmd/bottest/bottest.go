package bottest

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

const appName = `bottest`

var rootCmd = &cobra.Command{
	Use:   appName,
	Short: "Azure Bot Tester",
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("build version: %s\n", buildVersion)
	},
}

var (
	logLevel     = "info"
	buildVersion = "development"
)

func init() {
	rootCmd.AddCommand(versionCmd)

	rootCmd.PersistentFlags().StringVarP(&logLevel, "log.level", "l", "debug", "log level [debug, info, error]")
}

// Execute uses the args (os.Args[1:] by default)
// and run through the command tree finding appropriate matches
// for commands and then corresponding flags.
func Execute() {
	rootCmd.Execute()
}

func newLogger() *zap.SugaredLogger {
	var (
		cfg = zap.NewProductionConfig()
	)

	err := cfg.Level.UnmarshalText([]byte(logLevel))
	if err != nil {
		fmt.Println("error setting log level: ", err)
		os.Exit(1)
	}

	l, err := cfg.Build()
	if err != nil {
		fmt.Println("error creating logger: ", err)
		os.Exit(1)
	}

	return l.Sugar()
}
