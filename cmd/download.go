package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"time"

	"github.com/densestvoid/groupme"
	"github.com/manifoldco/promptui"
	"github.com/mmmveggies/groupme-files/internal/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Downloads images from a group",
	Run:   downloadCmdRun,
}

func init() {
	rootCmd.AddCommand(downloadCmd)
}

func downloadCmdRun(cmd *cobra.Command, args []string) {
	token := viper.GetString("token")
	log.Println("token", token)

	cwd, err := os.Getwd()
	util.IsOK(err)

	client := groupme.NewClient(token)
	groupId, err := util.PromptGroup(client)
	util.IsOK(err)

	base, err := (&promptui.Prompt{
		Label:   "Where should the files be downloaded to?",
		Default: path.Join(cwd, "groupme-downloads"),
	}).Run()
	util.IsOK(err)

	beginTimeStr, err := (&promptui.Prompt{
		Label:   "What is the upper age limit? (YYYY-MM-DD)",
		Default: time.Now().AddDate(0, -1, 0).Format("2006-01-02"),
	}).Run()
	util.IsOK(err)

	endTimeStr, err := (&promptui.Prompt{
		Label:   "What is the lower age limit? (YYYY-MM-DD)",
		Default: time.Now().Format("2006-01-02"),
	}).Run()
	util.IsOK(err)

	beginTime := util.MustBeValid(beginTimeStr)
	endTime := util.MustBeValid(endTimeStr)

	if beginTime.After(endTime) {
		log.Fatal("Beginning date cannot be after ending date")
	}

	err = downloadRange(client, base, groupId, beginTime, endTime)
	util.IsOK(err)
}

func downloadRange(
	client *groupme.Client,
	base string,
	groupId string,
	beginDate, endDate time.Time,
) error {
	ctx := context.Background()

	var beforeId groupme.ID

	log.Println("start", beginDate)
	log.Println("end", endDate)

	for {
		time.Sleep(time.Second)

		messages, err := client.IndexMessages(ctx, groupme.ID(groupId), &groupme.IndexMessagesQuery{
			BeforeID: beforeId,
			Limit:    100,
		})
		if err != nil {
			if strings.Contains(err.Error(), "304") {
				return nil
			}
			return err
		}

		ms := messages.Messages
		beforeId = ms[len(ms)-1].ID
		log.Printf("Reading page starting at: %s", ms[0].CreatedAt.ToTime())

		for _, m := range messages.Messages {
			ts := m.CreatedAt.ToTime()
			if ts.Before(beginDate) {
				return nil
			}
			if ts.After(endDate) {
				continue
			}

			userpath := path.Join(base, strings.Split(m.Name, " ")[0])

			for i, a := range m.Attachments {
				if a.Type == groupme.Image {
					util.IsOK(os.MkdirAll(userpath, os.ModePerm))

					bits := strings.Split(a.URL, ".")
					ext := bits[len(bits)-2]

					name := fmt.Sprintf("%s__%02d.%s", ts.Format("2006-01-02_15-04-05"), i, ext)
					filepath := path.Join(userpath, name)

					if _, err := os.Stat(filepath); err == nil {
						log.Printf("File already exists: %s", filepath)
						continue
					}
					if err := util.DownloadFile(a.URL, filepath); err != nil {
						return err
					}
					log.Printf("Downloaded: %s", filepath)
				}
			}
		}
	}
}
