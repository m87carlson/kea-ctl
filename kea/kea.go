package kea

type Command struct {
	Command string   `json:"command"`
	Service []string `json:"service"`
}

type HeartBeatArguments struct {
	DateTime string `json:"date-time"`
	State    string `json:"state"`
}

/*
type 	KeaResults []struct {
	Arguments HeartBeatArguments `json:"arguments"`
	Result int    `json:"result"`
	Text   string `json:"text"`
}
*/

type Results []struct {
	Arguments struct {
		DateTime string `json:"date-time"`
		State    string `json:"state"`
	} `json:"arguments"`
	Result int    `json:"result"`
	Text   string `json:"text"`
}

type VersionResults []struct {
	Arguments struct {
		Extended string `json:"extended"`
	} `json:"arguments"`
	Result int    `json:"result"`
	Text   string `json:"text"`
}
