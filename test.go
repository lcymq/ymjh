package server

type CampaignData struct {
}

type CampaignConfig struct {
	Dict         map[int]CampaignData
	MaxLevelDict map[int]int
	EndFlagDict  map[int]int

	lastCurLevelCount int
	levelCount        int
	settingIndex      int
}

func (c *CampaignConfig) CampaignConfig() {
	c.Dict = map[int]CampaignData{}
}
