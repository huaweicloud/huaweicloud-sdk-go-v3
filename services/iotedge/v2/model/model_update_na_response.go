package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateNaResponse struct {

	// NA系统ID，提供给其他系统访问的唯一标识
	NaId *string `json:"na_id,omitempty" xml:"na_id"`

	// NA系统名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 北向NA系统描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 访问URL地址
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint"`

	// 鉴权方式
	AuthType *string `json:"auth_type,omitempty" xml:"auth_type"`

	// 接入类型
	AccessType *string `json:"access_type,omitempty" xml:"access_type"`

	AccessRomaInfo *AccessRomaBriefInfo `json:"access_roma_info,omitempty" xml:"access_roma_info"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty" xml:"create_time"`

	// 更新时间
	UpdateTime     *string `json:"update_time,omitempty" xml:"update_time"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateNaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNaResponse struct{}"
	}

	return strings.Join([]string{"UpdateNaResponse", string(data)}, " ")
}
