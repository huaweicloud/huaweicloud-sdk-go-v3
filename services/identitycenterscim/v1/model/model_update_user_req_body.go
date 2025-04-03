package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateUserReqBody struct {

	// 用户的全局唯一标识符（ID）
	Id *string `json:"id,omitempty"`

	// 外部标识符
	ExternalId *string `json:"externalId,omitempty"`

	// 用户名，用于标识用户的唯一字符串
	UserName string `json:"userName"`

	Name *NameDto `json:"name"`

	// 包含用户显示名称的字符串
	DisplayName string `json:"displayName"`

	// 包含用户昵称的字符串
	NickName *string `json:"nickName,omitempty"`

	// 包含可能与用户关联的URL的字符串
	ProfileUrl *string `json:"profileUrl,omitempty"`

	// 包含用户的电子邮箱信息的对象列表
	Emails []EmailItemDto `json:"emails"`

	// 包含用户地址信息的对象列表
	Addresses *[]AddressItemDto `json:"addresses,omitempty"`

	// 包含用户电话号码信息的对象列表
	PhoneNumbers *[]PhoneNumberItemDto `json:"phoneNumbers,omitempty"`

	// 指示用户类型的字符串
	UserType *string `json:"userType,omitempty"`

	// 包含用户头衔的字符串
	Title *string `json:"title,omitempty"`

	// 包含用户首选语言的字符串
	PreferredLanguage *string `json:"preferredLanguage,omitempty"`

	// 包含用户地理区域或位置的字符串
	Locale *string `json:"locale,omitempty"`

	// 包含用户时区的字符串
	Timezone *string `json:"timezone,omitempty"`

	// 表示用户是否启用
	Active *bool `json:"active,omitempty"`

	// 概要
	Schemas []string `json:"schemas"`

	Urnietfparamsscimschemasextensionenterprise20User *EnterpriseDto `json:"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User,omitempty"`
}

func (o UpdateUserReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserReqBody struct{}"
	}

	return strings.Join([]string{"UpdateUserReqBody", string(data)}, " ")
}
