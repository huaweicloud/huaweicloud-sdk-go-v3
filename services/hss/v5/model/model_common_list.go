package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CommonList struct {

	// 提示信息列表
	DataList *[]string `json:"data_list,omitempty"`

	// 提示信息总数
	TotalNum *int32 `json:"total_num,omitempty"`
}

func (o CommonList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommonList struct{}"
	}

	return strings.Join([]string{"CommonList", string(data)}, " ")
}
