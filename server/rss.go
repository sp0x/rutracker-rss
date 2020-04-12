package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/sp0x/rutracker-rss/db"
	"github.com/sp0x/rutracker-rss/server/rss"
	"github.com/sp0x/rutracker-rss/torrent"
	"net/url"
)

func (s *Server) serveMusic(c *gin.Context) {
	torrents := storage.GetTorrentsInCategories([]int{
		409,  // Classical and modern academic music
		1125, // Folklore, national and ethnical music
		1849, //New age, relax, meditative and flamenco
		408,  //Rap, hip-hop and rnb
		1760, //Raggae, ska, dub
		416,  // OST, karaoke and musicals
		413,  //Other music
		2497, //Popular foreign music
	})
	rss.SendRssFeed(hostname, "music", torrents, c)
}

func (s *Server) serveAnime(c *gin.Context) {
	torrents := storage.GetTorrentsInCategories([]int{
		33, // Anime
	})
	rss.SendRssFeed(hostname, "anime", torrents, c)
}

func (s *Server) searchAndServe(c *gin.Context) {
	var search *torrent.Search
	ops := s.tracker.GetDefaultOptions()
	currentPage := uint(0)
	name := c.Param("name")
	name = url.QueryEscape(name)
	var items []db.Torrent
	for true {
		var err error
		if search == nil {
			search, err = s.tracker.Search(nil, name, 0)
		} else {
			search, err = s.tracker.Search(search, name, currentPage)
		}
		if err != nil {
			log.Warningf("Error while searching for torrent: %s . %s", name, err)
		}
		if currentPage >= ops.PageCount {
			break
		}
		s.tracker.ParseTorrents(search.GetDocument(), func(i int, tr *db.Torrent) {
			isNew, isUpdate := torrent.HandleTorrentDiscovery(s.tracker, tr)
			if isNew || isUpdate {
				if isNew && !isUpdate {
					_, _ = fmt.Fprintf(s.tabWriter, "Found new torrent #%s:\t%s\t[%s]:\t%s\n",
						tr.TorrentId, tr.AddedOnStr(), tr.Fingerprint, tr.Name)
				} else {
					_, _ = fmt.Fprintf(s.tabWriter, "Updated torrent #%s:\t%s\t[%s]:\t%s\n",
						tr.TorrentId, tr.AddedOnStr(), tr.Fingerprint, tr.Name)
				}
			} else {
				_, _ = fmt.Fprintf(s.tabWriter, "Torrent #%s:\t%s\t[%s]:\t%s\n",
					tr.TorrentId, tr.AddedOnStr(), "#", tr.Name)
			}
			items = append(items, *tr)
			s.tabWriter.Flush()
		})

		currentPage++
	}
	rss.SendRssFeed(hostname, name, items, c)
}

func (s *Server) serveShows(c *gin.Context) {
	torrents := storage.GetTorrentsInCategories([]int{
		189,  //Foreign shows
		2366, //Foreign shows in HD
		2100, //Asian shows
	})
	rss.SendRssFeed(hostname, "shows", torrents, c)
}

func (s *Server) serveMovies(c *gin.Context) {
	torrents := storage.GetTorrentsInCategories([]int{
		7,    //foreign films
		124,  //art-house and author movies
		93,   //DVD
		2198, //HD Video
		4,    //Multifilms
		352,  //3d/stereo movies, video, tv and sports
	})
	rss.SendRssFeed(hostname, "movies", torrents, c)
}