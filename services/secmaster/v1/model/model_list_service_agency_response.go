package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServiceAgencyResponse Response Object
type ListServiceAgencyResponse struct {

	// 当前账号是否已创建云服务关联委托
	CreateAgency *bool `json:"create_agency,omitempty"`

	// 角色列表
	RoleDescriptions *[]Role `json:"role_descriptions,omitempty"`
	HttpStatusCode   int     `json:"-"`
}

func (o ListServiceAgencyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServiceAgencyResponse struct{}"
	}

	return strings.Join([]string{"ListServiceAgencyResponse", string(data)}, " ")
}
