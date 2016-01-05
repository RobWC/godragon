package godragon

type Account struct {
	AccountID int
	PlatforID string
}

type Summoner struct {
	AccountID    int
	SummonerID   int
	PlatformID   string
	AllyBadge    int
	CurrentUser  Account
	ELO          map[string]int
	EnemeyBadge  int
	HighestRank  string
	SummonerName string
}

type SummonerList map[string]Summoner

func (c *APIClient) SummonerByName(name string) {

}
