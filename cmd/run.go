package cmd

import (
	"github.com/aliakbariaa1996/URL-Shortening/config"
	"github.com/aliakbariaa1996/URL-Shortening/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"os"
)

func init() {

	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	logrus.SetLevel(logrus.InfoLevel)
}

// Server is exported to make it graceful stop inside the main:
// cmd.Server.GracefulStop() for reconfiguration and code profiling.
var Server *grpc.Server

var runCMD = &cobra.Command{
	Use:   "run",
	Short: "Run totem user profile",
	Long:  `Run totem user profile`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		cmd.Flags().String("port", "5050", "HTTP server listen address")
		cmd.Flags().String("config", "", "config file if present")
		//redis flags
		cmd.Flags().String("db.redis_user", "admin", "Define redis user")
		cmd.Flags().String("db.redis_pass", "admin", "Define redis db password")
		cmd.Flags().String("db.redis_db", "admin", "Define redis db name")
		cmd.Flags().String("db.redis_host", "localhost", "Define redis host address .e.g localhost")
		cmd.Flags().Int("db.redis_port", 6379, "Define redis host address .e.g localhost")

		err := cmd.ParseFlags(args)
		if err != nil {
			return err
		}
		configFlag := cmd.Flags().Lookup("config")
		if configFlag != nil {
			configFilePath := configFlag.Value.String()
			if configFilePath != "" {
				viper.SetConfigFile(configFilePath)
				err := viper.ReadInConfig()
				if err != nil {
					return err
				}
			}
		}
		err = viper.BindPFlags(cmd.Flags())
		if err != nil {
			return err
		}
		return nil
	},
	RunE: runCmdE,
}

func intConfig() *config.Config {
	return &config.Config{
		Port:       viper.GetString("port"),
		DBHost:     viper.GetString("db.redis_host"),
		DBPassword: viper.GetString("db.redis_pass"),
		DBUser:     viper.GetString("db.redis_user"),
		DBPort:     viper.Get("db.redis_port").(int),
	}
}

func runCmdE(cmd *cobra.Command, args []string) error {
	cfg := intConfig()
	logger := logrus.New()

	if err := server.RunServer(cfg, logger); err != nil {
		log.Fatalf("%s", err.Error())
		return err
	}
	return nil
}

func init() {
	RootCmd.AddCommand(runCMD)
}
