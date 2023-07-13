package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAppCountResponse Response Object
type ShowAppCountResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 已使用数
	Used           *int32 `json:"used,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowAppCountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAppCountResponse struct{}"
	}

	return strings.Join([]string{"ShowAppCountResponse", string(data)}, " ")
}
