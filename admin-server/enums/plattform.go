package enums

type Plattform uint8

const (
	PlattformTwitch Plattform = iota
	PlattformYouTube
)

var plattformMap map[Plattform]string = map[Plattform]string{
	PlattformTwitch:  "Twitch",
	PlattformYouTube: "YouTube",
}

func (p Plattform) ToString() string {
	return plattformMap[p]
}
