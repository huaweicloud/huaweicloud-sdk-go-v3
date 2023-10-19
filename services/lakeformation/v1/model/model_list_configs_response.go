package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConfigsResponse Response Object
type ListConfigsResponse struct {

	// 配置项
	Configs *[]LakeCatConfiguration `json:"configs,omitempty"`

	PageInfo       *PagedInfo `json:"page_info,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigsResponse struct{}"
	}

	return strings.Join([]string{"ListConfigsResponse", string(data)}, " ")
}
