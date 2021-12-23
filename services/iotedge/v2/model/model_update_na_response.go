package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateNaResponse struct {
	// NA系统ID，提供给其他系统访问的唯一标识

	NaId *string `json:"na_id,omitempty"`
	// NA系统名称

	Name *string `json:"name,omitempty"`
	// 北向NA系统描述

	Description *string `json:"description,omitempty"`
	// 访问URL地址

	Endpoint *string `json:"endpoint,omitempty"`
	// 鉴权方式

	AuthType *string `json:"auth_type,omitempty"`
	// 接入类型

	AccessType *string `json:"access_type,omitempty"`

	AccessRomaInfo *AccessRomaBriefInfo `json:"access_roma_info,omitempty"`
	// 创建时间

	CreateTime *string `json:"create_time,omitempty"`
	// 更新时间

	UpdateTime     *string `json:"update_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateNaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNaResponse struct{}"
	}

	return strings.Join([]string{"UpdateNaResponse", string(data)}, " ")
}
