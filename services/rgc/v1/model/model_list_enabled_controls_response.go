package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEnabledControlsResponse Response Object
type ListEnabledControlsResponse struct {

	// 开启控制策略信息。
	EnabledControls *[]EnabledControl `json:"enabled_controls,omitempty"`

	PageInfo       *PageInfoDto `json:"page_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListEnabledControlsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnabledControlsResponse struct{}"
	}

	return strings.Join([]string{"ListEnabledControlsResponse", string(data)}, " ")
}
