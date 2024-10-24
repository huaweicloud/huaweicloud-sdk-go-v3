package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAccountsForProvisionedPermissionSetResponse Response Object
type ListAccountsForProvisionedPermissionSetResponse struct {

	// 账号ID列表
	AccountIds *[]string `json:"account_ids,omitempty"`

	PageInfo       *PageInfoDto `json:"page_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListAccountsForProvisionedPermissionSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccountsForProvisionedPermissionSetResponse struct{}"
	}

	return strings.Join([]string{"ListAccountsForProvisionedPermissionSetResponse", string(data)}, " ")
}
