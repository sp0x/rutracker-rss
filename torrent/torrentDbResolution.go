package torrent

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/sp0x/rutracker-rss/db"
	"os"
	"text/tabwriter"
)

//Gets torrent information from a given tracker and updates the torrent db
func ResolveTorrents(client *Rutracker, hours int) {
	gdb := db.GetOrmDb()
	defer gdb.Close()
	torrents := client.storage.GetOlderThanHours(hours)
	tabWr := new(tabwriter.Writer)
	tabWr.Init(os.Stdout, 0, 8, 0, '\t', 0)

	for i, t := range torrents {
		//Skip already resolved torrents.
		if t.Announce != "" {
			continue
		}
		def, err := ParseTorrentFromUrl(client, t.DownloadLink)
		if err != nil {
			log.Debugf("Could not resolve torrent: [%v] %v", t.TorrentId, t.Name)
			continue
		}
		t.Announce = def.Announce
		t.Publisher = def.Publisher
		t.AltName = def.Info.Name
		t.Size = def.GetTotalFileSize()
		perc := (float32(i) / float32(len(torrents))) * 100
		_, _ = fmt.Fprintf(tabWr, "%f%% Resolved [%s]\t%s\n", perc, t.TorrentId, t.Name)
		gdb.Save(t)
		_ = tabWr.Flush()
	}
}
