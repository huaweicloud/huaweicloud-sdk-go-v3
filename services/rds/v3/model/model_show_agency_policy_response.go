package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAgencyPolicyResponse Response Object
type ShowAgencyPolicyResponse struct {

	// 委托是否存在。
	IsExisted *bool `json:"is_existed,omitempty"`

	// 委托名称。
	Name *string `json:"name,omitempty"`

	// 委托角色列表。
	Roles          *[]AgencyRole `json:"roles,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowAgencyPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAgencyPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowAgencyPolicyResponse", string(data)}, " ")
}
