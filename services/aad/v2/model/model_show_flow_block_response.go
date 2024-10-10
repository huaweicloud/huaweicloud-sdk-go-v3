package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFlowBlockResponse Response Object
type ShowFlowBlockResponse struct {

	// total
	Total *int32 `json:"total,omitempty"`

	// ips
	Ips            *[]Ips `json:"ips,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowFlowBlockResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlowBlockResponse struct{}"
	}

	return strings.Join([]string{"ShowFlowBlockResponse", string(data)}, " ")
}
