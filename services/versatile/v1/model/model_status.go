package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Status struct {

	// 错误码
	Code int32 `json:"code"`

	// 错误描述信息
	Desc string `json:"desc"`
}

func (o Status) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Status struct{}"
	}

	return strings.Join([]string{"Status", string(data)}, " ")
}
