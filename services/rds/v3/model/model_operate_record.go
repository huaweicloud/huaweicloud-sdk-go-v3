package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OperateRecord struct {

	// 操作类型
	OperateType *string `json:"operate_type,omitempty"`

	// 用户名称
	UserName *string `json:"user_name,omitempty"`

	// 操作时间
	OperateTime *int64 `json:"operate_time,omitempty"`

	// 事件等级
	Level *string `json:"level,omitempty"`
}

func (o OperateRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperateRecord struct{}"
	}

	return strings.Join([]string{"OperateRecord", string(data)}, " ")
}
