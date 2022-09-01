package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateNaRequestDto struct {

	// NA系统名称
	Name string `json:"name" xml:"name"`

	// 北向NA系统描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 访问URL地址
	Endpoint string `json:"endpoint" xml:"endpoint"`

	// 鉴权方式
	AuthType *string `json:"auth_type,omitempty" xml:"auth_type"`

	AuthAkskInfo *AuthAkSkInfo `json:"auth_aksk_info,omitempty" xml:"auth_aksk_info"`

	// 接入类型
	AccessType string `json:"access_type" xml:"access_type"`

	AccessRomaInfo *AccessRomaInfo `json:"access_roma_info,omitempty" xml:"access_roma_info"`
}

func (o UpdateNaRequestDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNaRequestDto struct{}"
	}

	return strings.Join([]string{"UpdateNaRequestDto", string(data)}, " ")
}
