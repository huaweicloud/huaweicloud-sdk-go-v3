package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUsersRequest Request Object
type ListUsersRequest struct {

	// LakeFormation实例ID。创建实例时自动生成。例如：2180518f-42b8-4947-b20b-adfc53981a25。
	InstanceId string `json:"instance_id"`

	// 查询的用户来源。只能为IAM或SAML或LDAP或LOCAL或AGENTTENANT或OTHER。默认为IAM。
	UserSource *string `json:"user_source,omitempty"`

	// 查询返回条数。默认值为1000。最小值为1，最大值为2000。
	Limit *int32 `json:"limit,omitempty"`

	// 查询的起始记录ID。最小长度为0，最大长度为256。
	Marker *string `json:"marker,omitempty"`

	// 是否查询上一页。默认为false。
	ReversePage *bool `json:"reverse_page,omitempty"`

	// 用户模糊查询。只能包含字母、数字和_|*.特殊字符，且长度为1~255个字符。
	UserNamePattern *string `json:"user_name_pattern,omitempty"`
}

func (o ListUsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUsersRequest struct{}"
	}

	return strings.Join([]string{"ListUsersRequest", string(data)}, " ")
}
