package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAgencyInfoResponse Response Object
type ShowAgencyInfoResponse struct {

	// 委托是否存在。
	IsExisted *bool `json:"is_existed,omitempty"`

	// 委托名称。
	Name *string `json:"name,omitempty"`

	// 委托绑定的权限策略信息。
	Roles          *[]AgencyRole `json:"roles,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowAgencyInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAgencyInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowAgencyInfoResponse", string(data)}, " ")
}
