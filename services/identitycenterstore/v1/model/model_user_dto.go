package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserDto A user object that contains a specified user’s metadata and attributes.
type UserDto struct {

	// 用户的地址信息列表
	Addresses *[]AddressDto `json:"addresses,omitempty"`

	// 用户的显示名称
	DisplayName string `json:"display_name"`

	// 用户的电子邮箱信息列表
	Emails []EmailDto `json:"emails"`

	// 外部身份源分配给此资源的标识符
	ExternalId *string `json:"external_id,omitempty"`

	// 用户的外部标识符信息列表
	ExternalIds *[]ExternalIdDto `json:"external_ids,omitempty"`

	// 身份源的全局唯一标识符（ID）
	IdentityStoreId string `json:"identity_store_id"`

	// 用户的地理区域或位置信息
	Locale *string `json:"locale,omitempty"`

	Name *NameDto `json:"name"`

	// 用户昵称
	Nickname *string `json:"nickname,omitempty"`

	// 用户的电话号码信息列表
	PhoneNumbers *[]PhoneNumberDto `json:"phone_numbers,omitempty"`

	// 用户语言首选项
	PreferredLanguage *string `json:"preferred_language,omitempty"`

	// 与用户关联的URL
	ProfileUrl *string `json:"profile_url,omitempty"`

	// 用户时区
	Timezone *string `json:"timezone,omitempty"`

	// 用户头衔
	Title *string `json:"title,omitempty"`

	// 身份源中IAM身份中心用户的全局唯一标识符（ID）
	UserId string `json:"user_id"`

	// 用户名，用于标识用户的唯一字符串
	UserName string `json:"user_name"`

	// 用户类型
	UserType *string `json:"user_type,omitempty"`

	// 创建用户时的时间戳
	CreatedAt int64 `json:"created_at"`

	// 创建者
	CreatedBy string `json:"created_by"`

	// 更新用户时的时间戳
	UpdatedAt int64 `json:"updated_at"`

	// 更新者
	UpdatedBy string `json:"updated_by"`

	// 一个布尔值，表示用户是否启用
	Enabled bool `json:"enabled"`

	Enterprise *EnterpriseDto `json:"enterprise,omitempty"`
}

func (o UserDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserDto struct{}"
	}

	return strings.Join([]string{"UserDto", string(data)}, " ")
}
