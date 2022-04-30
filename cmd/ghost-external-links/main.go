package main

import (
	"fmt"
	"github.com/lfritz/env"
	log "github.com/sirupsen/logrus"
	"github.com/sklinkert/ghost"
	"regexp"
	"sort"
	"strings"
)

func main() {
	var url, contentKey, adminKey string
	var e = env.New()
	e.String("GHOST_ADMIN_KEY", &adminKey, "Ghost: Admin API Key")
	e.String("GHOST_CONTENT_KEY", &contentKey, "Ghost: Content API Key")
	e.String("GHOST_URL", &url, "Ghost: Blog URL")
	if err := e.Load(); err != nil {
		log.WithError(err).Fatal("env loading failed")
	}

	if url == "" {
		log.Fatal("GHOST_URL is required")
	}
	if contentKey == "" {
		log.Fatal("GHOST_CONTENT_KEY is required")
	}
	if adminKey == "" {
		log.Fatal("GHOST_ADMIN_KEY is required")
	}

	ghostCli, err := ghost.New(url, contentKey, adminKey)
	if err != nil {
		log.WithError(err).Fatal("ghost client creation failed")
	}
	posts, err := ghostCli.GetPosts()
	if err != nil {
		log.WithError(err).Fatal("ghost posts retrieval failed")
	}
	pages, err := ghostCli.GetPages()
	if err != nil {
		log.WithError(err).Fatal("ghost pages retrieval failed")
	}

	var urlCount = map[string]int{} // link url -> count
	var urls = map[string]string{}  // link url -> page/post url
	var urlRegexp = regexp.MustCompile(`href="(https?://[^"]+)"`)
	for _, post := range posts.Posts {
		for _, match := range urlRegexp.FindAllStringSubmatch(post.HTML, -1) {
			link := match[1]
			if !strings.HasPrefix(link, url) {
				urls[link] = post.URL
				urlCount[link]++
			}
		}
	}
	for _, page := range pages.Pages {
		for _, match := range urlRegexp.FindAllStringSubmatch(page.HTML, -1) {
			link := match[1]
			if !strings.HasPrefix(link, url) {
				urls[link] = page.URL
				urlCount[link]++
			}
		}
	}

	type kv struct {
		Key   string
		Value int
	}

	var ss []kv
	for k, v := range urlCount {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value < ss[j].Value
	})

	var totalCount = 0
	for _, kv := range ss {
		fmt.Printf("%5d | %s\n\tSource: %s\n\n", kv.Value, kv.Key, urls[kv.Key])
		totalCount += kv.Value
	}
	log.Infof("Links found: %d", totalCount)
}
