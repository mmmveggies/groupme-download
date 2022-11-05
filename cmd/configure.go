package cmd

import (
	// "errors"
	"fmt"
	"path"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/mmmveggies/groupme-files/internal/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configures your API Token",
	Run:   configureCmdRun,
}

var tokenPrompt = promptui.Prompt{
	Label: "Access Token",
}

func init() {
	rootCmd.AddCommand(configureCmd)
}

func configureCmdRun(cmd *cobra.Command, args []string) {
	raw, err := tokenPrompt.Run()
	util.IsOK(err)
	token := strings.TrimSpace(raw)

	viper.Set("token", token)

	appDir, err := util.GetApplicationDir()
	util.IsOK(err)

	err = viper.WriteConfigAs(path.Join(appDir, ".groupme-files.env"))
	util.IsOK(err)

	fmt.Printf("Updated configuration in: %s\n", viper.ConfigFileUsed())
}
