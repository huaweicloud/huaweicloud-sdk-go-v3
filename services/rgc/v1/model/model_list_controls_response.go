package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListControlsResponse Response Object
type ListControlsResponse struct {

	// 控制策略信息。
	Controls *[]Control `json:"controls,omitempty"`

	PageInfo       *PageInfoDto `json:"page_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListControlsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListControlsResponse struct{}"
	}

	return strings.Join([]string{"ListControlsResponse", string(data)}, " ")
}
