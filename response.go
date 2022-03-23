package main

type Song struct {
	Id     string
	Text   string
	Artist string
	Title  string
	Album  string
	Genre  string
	Art    string
}

type SongPlayer struct {
	PlayedAt int64 `json:"played_at"`
	Song     Song
}

type AzuraResponse struct {
	NowPlaying  SongPlayer `json:"now_playing"`
	PlayingNext SongPlayer `json:"playing_next"`
}
