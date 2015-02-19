package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
)

var maxItemLength = 25

func main() {
	doc, err := goquery.NewDocument("http://lyrics.wikia.com/Kidz_Bop")
	if err != nil {
		panic(err)
	}
	doc.Find("#mw-content-text").Find("h2").Each(getAlbums)
}

func getAlbums(i int, s *goquery.Selection) {
	album := s.Find("a").First()
	album_name := strings.Split(album.Text(), " ")
	if album_name[0] == "Kidz" {
		_, err := strconv.Atoi(album_name[2])
		if album_name[2][0] == '(' || err == nil {
			album_name_fixed := strings.Join(album_name[0:len(album_name)-1], " ")
			fmt.Printf("Found: %s\n", album_name_fixed)
			nextElem := s.Next()
			for !nextElem.Is("h2") {
				if nextElem.Is("ol") {
					nextElem.Find("li").Each(getSongs)
				}
				nextElem = nextElem.Next()
			}
		}
	}
}
func getSongs(i int, s *goquery.Selection) {
	song := s.Find("a").First()
	fmt.Printf("  Found: %s\n", song.Text())
}
