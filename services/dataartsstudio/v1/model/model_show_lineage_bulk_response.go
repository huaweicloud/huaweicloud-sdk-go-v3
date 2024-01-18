package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLineageBulkResponse Response Object
type ShowLineageBulkResponse struct {

	// 作业算子个数，批量查询根据作业算子获取血缘
	Total float32 `json:"total,omitempty"`

	// 当前页作业算子包含的表血缘列表
	Data           *[]TableLineageInfo `json:"data,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowLineageBulkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLineageBulkResponse struct{}"
	}

	return strings.Join([]string{"ShowLineageBulkResponse", string(data)}, " ")
}
