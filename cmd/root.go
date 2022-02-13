package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/mmmveggies/groupme-files/internal/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "groupme-files",
		Short: "Download GroupMe media.",
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := util.GetApplicationDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("env")
		viper.SetConfigName(".groupme-files")
		cfgFile = path.Join(home, ".groupme-files.env")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		_, err = os.Create(cfgFile)
		util.IsOK(err)
	}
	util.IsOK(viper.ReadInConfig())

	fmt.Println("Using config file:", viper.ConfigFileUsed())
}
