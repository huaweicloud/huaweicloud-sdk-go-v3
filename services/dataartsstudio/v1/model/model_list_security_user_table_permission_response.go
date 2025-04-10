package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityUserTablePermissionResponse Response Object
type ListSecurityUserTablePermissionResponse struct {

	// 需要查询的表
	TableList *[]SecurityListUserTableListTableList `json:"table_list,omitempty"`

	Proposer       *SecurityListUserTableListProposer `json:"proposer,omitempty"`
	HttpStatusCode int                                `json:"-"`
}

func (o ListSecurityUserTablePermissionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityUserTablePermissionResponse struct{}"
	}

	return strings.Join([]string{"ListSecurityUserTablePermissionResponse", string(data)}, " ")
}
