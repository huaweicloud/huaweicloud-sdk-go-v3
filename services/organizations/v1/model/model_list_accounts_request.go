package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAccountsRequest Request Object
type ListAccountsRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// 父节点（根或组织单元）的唯一标识符（ID）。
	ParentId *string `json:"parent_id,omitempty"`

	// 是否返回账号邮箱、手机号信息。若此参数为True，Limit最多200。
	WithRegisterContactInfo *bool `json:"with_register_contact_info,omitempty"`

	// 页面中最大结果数量。
	Limit *int32 `json:"limit,omitempty"`

	// 分页标记。
	Marker *string `json:"marker,omitempty"`
}

func (o ListAccountsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccountsRequest struct{}"
	}

	return strings.Join([]string{"ListAccountsRequest", string(data)}, " ")
}
