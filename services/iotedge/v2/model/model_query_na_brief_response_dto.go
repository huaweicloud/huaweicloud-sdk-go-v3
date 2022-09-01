package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建及查询北向NA返回结构体（简洁版）
type QueryNaBriefResponseDto struct {

	// NA系统ID，提供给其他系统访问的唯一标识
	NaId *string `json:"na_id,omitempty" xml:"na_id"`

	// NA系统名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 北向NA系统描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 访问URL地址
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty" xml:"create_time"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty" xml:"update_time"`
}

func (o QueryNaBriefResponseDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryNaBriefResponseDto struct{}"
	}

	return strings.Join([]string{"QueryNaBriefResponseDto", string(data)}, " ")
}
