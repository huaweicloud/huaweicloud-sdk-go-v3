package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SaveImChannelsResponse Response Object
type SaveImChannelsResponse struct {

	// 配置 ID 列表
	ConfigIds      *[]string `json:"config_ids,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o SaveImChannelsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveImChannelsResponse struct{}"
	}

	return strings.Join([]string{"SaveImChannelsResponse", string(data)}, " ")
}
