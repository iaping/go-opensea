package opensea

type Offer struct {
	ItemType             int    `json:"itemType"`
	Token                string `json:"token"`
	IdentifierOrCriteria string `json:"identifierOrCriteria"`
	StartAmount          string `json:"startAmount"`
	EndAmount            string `json:"endAmount"`
}

type Consideration struct {
	Offer
	Recipient string `json:"recipient"`
}

type ProtocolData struct {
	Parameters struct {
		Offerer                         string          `json:"offerer"`
		Offer                           []Offer         `json:"offer"`
		Consideration                   []Consideration `json:"consideration"`
		StartTime                       string          `json:"startTime"`
		EndTime                         string          `json:"endTime"`
		OrderType                       int             `json:"orderType"`
		Zone                            string          `json:"zone"`
		ZoneHash                        string          `json:"zoneHash"`
		Salt                            string          `json:"salt"`
		ConduitKey                      string          `json:"conduitKey"`
		TotalOriginalConsiderationItems int             `json:"totalOriginalConsiderationItems"`
		Counter                         int             `json:"counter"`
	} `json:"parameters"`
	Signature string `json:"signature"`
}

type Account struct {
	Id            int64  `json:"user"`
	ProfileImgUrl string `json:"profile_img_url"`
	Address       string `json:"address"`
	Config        string `json:"config"`
}
