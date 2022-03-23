package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

var sleepTime int = 30
var telosSuccess bool = true
var tweetSuccess bool = true
var saveSuccess bool = true

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	telosHost := os.Getenv("TELOS_HOST")
	telosPort := os.Getenv("TELOS_PORT")
	azuraEndpoint := os.Getenv("AZURACAST_URL")

	// Fallback environment variables
	if telosHost == "" {
		telosHost = "127.0.0.1"
	}
	if telosPort == "" {
		telosPort = "9090"
	}
	if azuraEndpoint == "" {
		azuraEndpoint = "https://radio.mdogradio.com/api/nowplaying_static/non_processed.json"
	}

	// Create Twitter client
	twitterClient, err := createTwitterClient()
	if err != nil {
		log.Println(err)
		time.Sleep(5 * time.Second)
	}

	// Load last played song from file
	title, artist := loadLastPlayed()
	lastTitle, lastArtist := loadLastPlayed()
	tweetedTitle, tweetedArtist := loadLastPlayed()

	// Keep program running in the event of a recoverable error
	for {
		// Create TCP client
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", telosHost, telosPort))
		if err != nil {
			log.Println(err)
			time.Sleep(5 * time.Second)
			continue
		}

		// Connect to AzuraCast API
		for {
			var data AzuraResponse

			// Get data from API
			response, err := http.Get(azuraEndpoint)
			if err != nil {
				log.Println(err)
				time.Sleep(5 * time.Second)
				continue
			}

			// Handle unsuccessful request
			if response.StatusCode != 200 {
				log.Println("Error communicating with AzuraCast API")
				time.Sleep(5 * time.Second)
				continue
			}

			// Extract body from response
			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Println(err)
				time.Sleep(5 * time.Second)
				continue
			}

			// Unmarshal response body
			err = json.Unmarshal(body, &data)
			if err != nil {
				log.Println(err)
				time.Sleep(5 * time.Second)
				continue
			}

			// Repeat every second until a new API request should be made
			for i := 0; i < sleepTime; i++ {
				now := time.Now().Unix()

				// Update the current song
				if now < data.PlayingNext.PlayedAt {
					artist = data.NowPlaying.Song.Artist
					title = data.NowPlaying.Song.Title
				} else {
					artist = data.PlayingNext.Song.Artist
					title = data.PlayingNext.Song.Title
				}

				// Save current song to file
				err = saveLastPlayed(title, artist)
				if err != nil {
					saveSuccess = false
					log.Println(err)
				} else {
					saveSuccess = true
				}

				// Send data to Telos server
				if title != lastTitle || artist != lastArtist || !telosSuccess {
					_, err := fmt.Fprintf(conn, "<nowplaying><artist>%s</artist><title>%s</title><url>%s</url></nowplaying>", artist, title, "https://www.mdogradio.com/#listen")
					if err != nil {
						telosSuccess = false
						log.Println(err)
					} else {
						telosSuccess = true
						lastTitle = title
						lastArtist = artist
					}
				}

				// Post tweet
				if (title != tweetedTitle || artist != tweetedArtist || !tweetSuccess) && title != "" && artist != "" && saveSuccess {
					tweetText := fmt.Sprintf("Now playing: %s by %s", title, artist)
					err = postTweet(twitterClient, tweetText)
					if err != nil {
						tweetSuccess = false
						log.Println(err)
					} else {
						tweetSuccess = true
						tweetedTitle = title
						tweetedArtist = artist
					}
				}

				// Output to console
				if title != lastTitle || artist != lastArtist {
					log.Println("Now playing:", title, "by", artist)
				}

				// Wait 1 second
				time.Sleep(1 * time.Second)
			}
		}
	}
}
