package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUsersV5Request Request Object
type ListUsersV5Request struct {

	// 每页显示的条目数量，范围为1到200条，默认为100条。
	Limit *int32 `json:"limit,omitempty"`

	// 分页标记，长度为4到400个字符，只包含字母、数字、\"+\"、\"/\"、\"=\"、\"-\"和\"_\"的字符串。
	Marker *string `json:"marker,omitempty"`

	// 用户组ID，长度为1到64个字符，只包含字母、数字和\"-\"的字符串。
	GroupId *string `json:"group_id,omitempty"`

	// 永久访问密钥ID，即AK。
	AccessKeyId *string `json:"access_key_id,omitempty"`

	// **参数解释**： 服务专属凭证的标识符 ID。  **约束限制**： 不涉及。  **取值范围**： 长度固定为20，格式为HSSC + 16位随机字符，只能包含大小写字母和数字。  **默认取值**： 不涉及。
	ServiceSpecificCredentialId *string `json:"service_specific_credential_id,omitempty"`
}

func (o ListUsersV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUsersV5Request struct{}"
	}

	return strings.Join([]string{"ListUsersV5Request", string(data)}, " ")
}
