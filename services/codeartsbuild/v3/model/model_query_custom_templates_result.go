package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryCustomTemplatesResult struct {

	// 返回模板数量
	TotalSize *int32 `json:"total_size,omitempty"`

	// 模板信息列表
	Items *[]QueryTemplatesItems `json:"items,omitempty"`
}

func (o QueryCustomTemplatesResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryCustomTemplatesResult struct{}"
	}

	return strings.Join([]string{"QueryCustomTemplatesResult", string(data)}, " ")
}
