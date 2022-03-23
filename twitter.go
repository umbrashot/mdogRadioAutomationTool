package main

import (
	"context"
	"fmt"
	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweets"
	"github.com/michimani/gotwi/tweets/types"
	"log"
	"os"
	"strings"
)

func createTwitterClient() (*gotwi.GotwiClient, error) {
	in := &gotwi.NewGotwiClientInput{
		AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
		OAuthToken:           os.Getenv("TWITTER_USER_ACCESS_TOKEN"),
		OAuthTokenSecret:     os.Getenv("TWITTER_USER_ACCESS_TOKEN_SECRET"),
	}

	c, err := gotwi.NewGotwiClient(in)
	return c, err
}

func postTweet(c *gotwi.GotwiClient, m string) error {
	p := &types.ManageTweetsPostParams{
		Text: gotwi.String(m),
	}

	_, err := tweets.ManageTweetsPost(context.Background(), c, p)
	if err != nil {
		if strings.Contains(fmt.Sprint(err), "You are not allowed to create a Tweet with duplicate content.") {
			log.Println(err)
			return nil
		}
		return err
	}
	log.Printf("Tweeted:\"%s\"", m)
	return nil
}
