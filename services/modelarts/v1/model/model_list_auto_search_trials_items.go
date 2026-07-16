package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAutoSearchTrialsItems 超参搜索列表。
type ListAutoSearchTrialsItems struct {

	// 超参搜索所有trial结果的字段信息。
	Header *[]string `json:"header,omitempty"`

	// 超参搜索所有trial结果的每条数据列表。
	Data *[][]string `json:"data,omitempty"`
}

func (o ListAutoSearchTrialsItems) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAutoSearchTrialsItems struct{}"
	}

	return strings.Join([]string{"ListAutoSearchTrialsItems", string(data)}, " ")
}
