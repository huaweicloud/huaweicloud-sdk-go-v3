package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAgencyResponse Response Object
type ListAgencyResponse struct {

	// 委托名称
	AgencyName *string `json:"agency_name,omitempty"`

	// 授权与否。如果用户的服务委托权限包含系统所需要的委托权限返回true，否则返回false。
	Authorized *bool `json:"authorized,omitempty"`

	// 委托策略列表。
	AgencyPolicies *[]AgencyPolicy `json:"agency_policies,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListAgencyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAgencyResponse struct{}"
	}

	return strings.Join([]string{"ListAgencyResponse", string(data)}, " ")
}
