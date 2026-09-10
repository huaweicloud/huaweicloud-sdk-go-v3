package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOperateRecordResponse Response Object
type ListOperateRecordResponse struct {

	// 本次查询事件列表返回的事件记录的总条数
	Count *int32 `json:"count,omitempty"`

	// 本次查询事件列表返回的事件记录
	Traces *[]OperateRecord `json:"traces,omitempty"`

	// 所有事件类型
	AllOperateType *[]string `json:"all_operate_type,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListOperateRecordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOperateRecordResponse struct{}"
	}

	return strings.Join([]string{"ListOperateRecordResponse", string(data)}, " ")
}
