package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetUserResponse Response Object
type GetUserResponse struct {

	// 用户的全局唯一标识符（ID）
	Id *string `json:"id,omitempty"`

	// 外部标识符
	ExternalId *string `json:"externalId,omitempty"`

	Meta *MetaDto `json:"meta,omitempty"`

	// 概要
	Schemas *[]string `json:"schemas,omitempty"`

	// 用户名，用于标识用户的唯一字符串
	UserName *string `json:"userName,omitempty"`

	Name *NameDto `json:"name,omitempty"`

	// 包含用户显示名称的字符串
	DisplayName *string `json:"displayName,omitempty"`

	// 表示用户是否启用
	Active *bool `json:"active,omitempty"`

	// 包含用户的电子邮箱信息的对象列表
	Emails *[]EmailItemDto `json:"emails,omitempty"`

	// 指示用户类型的字符串
	UserType *string `json:"userType,omitempty"`

	// 用户昵称
	NickName *string `json:"nickName,omitempty"`

	// 用户头衔
	Title *string `json:"title,omitempty"`

	// 用户语言首选项
	PreferredLanguage *string `json:"preferredLanguage,omitempty"`

	// 用户的地理区域或位置信息
	Locale *string `json:"locale,omitempty"`

	// 用户时区
	Timezone *string `json:"timezone,omitempty"`

	// 用户的地址信息列表
	Addresses *[]AddressItemDto `json:"addresses,omitempty"`

	// 用户的电话号码信息列表
	PhoneNumbers *[]PhoneNumberItemDto `json:"phoneNumbers,omitempty"`

	Urnietfparamsscimschemasextensionenterprise20User *EnterpriseDto `json:"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User,omitempty"`
	HttpStatusCode                                    int            `json:"-"`
}

func (o GetUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetUserResponse struct{}"
	}

	return strings.Join([]string{"GetUserResponse", string(data)}, " ")
}
