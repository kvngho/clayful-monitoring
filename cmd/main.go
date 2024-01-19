package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kvngho/clayful-monitoring/cmd/command"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init () {
	rootCmd.AddCommand(command.CheckProductCmd)
}
var rootCmd = &cobra.Command{
	Use: "monitor [sub]",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		workDir, _ := os.Getwd()
		for {
			if _, err := os.Stat(filepath.Join(workDir, ".env")); err == nil {
				viper.SetConfigName(".env")
				viper.SetConfigType("dotenv")
				viper.AddConfigPath(workDir)
				if err := viper.ReadInConfig(); err != nil {
					panic(fmt.Errorf("failed to read config file: %w", err))
				}
				break
			}
			if workDir == "/" {
				break
			}
			workDir = filepath.Dir(workDir)
		}
		viper.AutomaticEnv()
		return nil
	},
}

func main() {
	rootCmd.Execute()
}
