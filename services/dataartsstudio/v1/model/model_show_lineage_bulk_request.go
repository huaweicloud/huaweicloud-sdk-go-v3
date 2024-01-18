package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLineageBulkRequest Request Object
type ShowLineageBulkRequest struct {

	// 实例id
	Instance string `json:"instance"`

	// 分页参数偏移量，默认值0
	Offset *int32 `json:"offset,omitempty"`

	// 分页参每页数量，默认值100
	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowLineageBulkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLineageBulkRequest struct{}"
	}

	return strings.Join([]string{"ShowLineageBulkRequest", string(data)}, " ")
}
