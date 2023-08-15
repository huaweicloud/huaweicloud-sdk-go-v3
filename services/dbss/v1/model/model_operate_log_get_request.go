package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OperateLogGetRequest struct {
	Time *TimeRangeBean `json:"time,omitempty"`

	// 筛选角色用户获取操作日志
	UserName *string `json:"user_name,omitempty"`

	// 筛选操作对象名称获取操作日志
	OperateName *string `json:"operate_name,omitempty"`

	// 根据执行结果获取操作日志 [success, fail]
	Result *string `json:"result,omitempty"`

	// 页数
	Page *string `json:"page,omitempty"`

	// 每页条数
	Size *string `json:"size,omitempty"`
}

func (o OperateLogGetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperateLogGetRequest struct{}"
	}

	return strings.Join([]string{"OperateLogGetRequest", string(data)}, " ")
}
