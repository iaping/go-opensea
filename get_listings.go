package opensea

import (
	"fmt"
)

type GetListings struct {
	Api
}

func NewGetListings(c *Client) *GetListings {
	api := &GetListings{}
	api.c = c
	return api
}

func (api *GetListings) Do(chain, protocol string, params *ListingsParam) (*ListingsResponse, error) {
	path := fmt.Sprintf("/api/v2/orders/%s/%s/listings", chain, protocol)
	req := api.request(path, false, params)
	resp := &ListingsResponse{}
	if err := api.c.Do(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type ListingsParam struct {
	AssetContractAddress string `url:"asset_contract_address,omitempty"`
	Bundled              bool   `url:"bundled,omitempty"`
	Cursor               string `url:"cursor,omitempty"`
	Limit                int    `url:"limit,omitempty"`
	ListedAfter          int64  `url:"listed_after,omitempty"`
	ListedBefore         int64  `url:"listed_before,omitempty"`
	Maker                string `url:"maker,omitempty"`
	OrderBy              string `url:"order_by,omitempty"`
	OrderDirection       string `url:"order_direction,omitempty"`
	PaymentTokenAddress  string `url:"payment_token_address,omitempty"`
	Taker                string `url:"taker,omitempty"`
	Token_ids            []int  `url:"token_ids,omitempty"`
}

type ListingsResponse struct {
	Next   string  `json:"next"`
	Orders []Order `json:"orders"`
}

type Order struct {
	CreatedDate       string        `json:"created_date"`
	ClosingDate       string        `json:"closing_date"`
	Listing_time      int64         `json:"listing_time"`
	ExpirationTime    int64         `json:"expiration_time"`
	OrderHash         string        `json:"order_hash"`
	ProtocolData      *ProtocolData `json:"protocol_data"`
	ProtocolAddress   string        `json:"protocol_address"`
	CurrentPrice      string        `json:"current_price"`
	Maker             *Account      `json:"maker"`
	Taker             *Account      `json:"taker"`
	Side              string        `json:"side"`
	OrderType         string        `json:"order_type"`
	Cancelled         bool          `json:"cancelled"`
	Finalized         bool          `json:"finalized"`
	MarkedInvalid     bool          `json:"marked_invalid"`
	RemainingQuantity int           `json:"remaining_quantity"`
	RelayId           string        `json:"relay_id"`
}
