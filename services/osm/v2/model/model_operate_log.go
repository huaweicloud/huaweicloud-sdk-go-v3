package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OperateLog struct {

	// 操作指令
	Oper *string `json:"oper,omitempty" xml:"oper"`

	// 操作时间
	OperateTime *string `json:"operate_time,omitempty" xml:"operate_time"`
}

func (o OperateLog) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperateLog struct{}"
	}

	return strings.Join([]string{"OperateLog", string(data)}, " ")
}
