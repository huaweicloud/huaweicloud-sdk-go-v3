package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDelegatedAdministratorsResponse Response Object
type ListDelegatedAdministratorsResponse struct {

	// 组织中委托管理员列表。
	DelegatedAdministrators *[]DelegatedAdministratorDto `json:"delegated_administrators,omitempty"`

	PageInfo       *PageInfoDto `json:"page_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListDelegatedAdministratorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDelegatedAdministratorsResponse struct{}"
	}

	return strings.Join([]string{"ListDelegatedAdministratorsResponse", string(data)}, " ")
}
