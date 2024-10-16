package modules


type Scrape interface {
	GetQuotes() (string,string)
}

type Quote struct {
	Text string
	Author string
}