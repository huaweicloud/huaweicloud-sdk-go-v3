package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryTemplatesResult struct {

	// 返回模板数量
	TotalSize *int32 `json:"total_size,omitempty"`

	// 模板信息列表
	Items *[]QueryTemplatesItems `json:"items,omitempty"`
}

func (o QueryTemplatesResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryTemplatesResult struct{}"
	}

	return strings.Join([]string{"QueryTemplatesResult", string(data)}, " ")
}
