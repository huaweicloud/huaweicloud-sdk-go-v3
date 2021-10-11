package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowNaResponse struct {
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

func (o ShowNaResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowNaResponse struct{}"
	}

	return strings.Join([]string{"ShowNaResponse", string(data)}, " ")
}
