package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建及查询北向NA返回结构体（简洁版）
type QueryNaBriefResponseDto struct {
	// NA系统ID，提供给其他系统访问的唯一标识

	NaId *string `json:"na_id,omitempty"`
	// NA系统名称

	Name *string `json:"name,omitempty"`
	// 北向NA系统描述

	Description *string `json:"description,omitempty"`
	// 访问URL地址

	Endpoint *string `json:"endpoint,omitempty"`
	// 创建时间

	CreateTime *string `json:"create_time,omitempty"`
	// 更新时间

	UpdateTime *string `json:"update_time,omitempty"`
}

func (o QueryNaBriefResponseDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryNaBriefResponseDto struct{}"
	}

	return strings.Join([]string{"QueryNaBriefResponseDto", string(data)}, " ")
}
