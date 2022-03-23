# Mdog Radio Automation Tool

This tool reads data from the AzuraCast API, feeds it into the Telos Z/IPStream software and tweets the currently playing song.

## Requirements

- Access to the Twitter API V2
- An OAuth 1.0a user access token with read/write permission for the Twitter API V2
- A running TCP server (Telos Z/IPStream) to accept XML song data
- A running AzuraCast server hosting the AzuraCast API

## Environment variables

You can set environment variables manually or set them in a `.env` file. An example `.env` file has been provided in this repo.

| Variable name                    | Purpose                                                       |
|----------------------------------|---------------------------------------------------------------|
| GOTWI_API_KEY                    | Your Twitter API consumer key                                 |  
| GOTWI_API_KEY_SECRET             | Your Twitter API consumer key secret                          |  
| TWITTER_USER_ACCESS_TOKEN        | The access token for your Twitter account                     | 
| TWITTER_USER_ACCESS_TOKEN_SECRET | The access token secret for your Twitter account              |  
| TELOS_HOST                       | The hostname / IP address for the Telos Z/IPStream TCP server | 
| TELOS_PORT                       | The network port for the Telos Z/IPStream TCP server          |
| AZURACAST_URL                    | The URL for the AzuraCast endpoint to retrieve data from      |
