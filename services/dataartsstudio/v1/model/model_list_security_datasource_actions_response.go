package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityDatasourceActionsResponse Response Object
type ListSecurityDatasourceActionsResponse struct {

	// 权限操作列表
	PermissionActions *[]PermissionActions `json:"permission_actions,omitempty"`
	HttpStatusCode    int                  `json:"-"`
}

func (o ListSecurityDatasourceActionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityDatasourceActionsResponse struct{}"
	}

	return strings.Join([]string{"ListSecurityDatasourceActionsResponse", string(data)}, " ")
}
