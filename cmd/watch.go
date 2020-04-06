package main

import (
	"fmt"
	"github.com/sp0x/rutracker-rss/rss"
	"github.com/sp0x/rutracker-rss/torrent"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var watchInterval int

func init() {
	cmdWatch := &cobra.Command{
		Use:   "watch",
		Short: "Watches the torrent tracker for new torrents.",
		Run:   watchTracker,
	}
	cmdWatch.Flags().IntVarP(&watchInterval, "interval", "i", 10, "Interval between checks.")
	viper.SetDefault("rss_port", 5000)
	_ = viper.BindEnv("rss_port")
	rootCmd.AddCommand(cmdWatch)
}

func watchTracker(cmd *cobra.Command, args []string) {
	client := torrent.NewRutracker()
	err := client.Login(viper.GetString("username"), viper.GetString("password"))
	if err != nil {
		fmt.Println("Could not login to tracker.")
		os.Exit(1)
	}
	go func() {
		err = rss.StartServer(viper.GetInt("rss_port"))
		if err != nil {
			fmt.Print(err)
		}
	}()

	torrent.Watch(client, watchInterval)
}
