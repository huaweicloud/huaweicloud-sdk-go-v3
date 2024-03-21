package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityDatasourceConfigurationsResponse Response Object
type ListSecurityDatasourceConfigurationsResponse struct {

	// 数据源操作权限列表
	Configurations *[]PermissionConfiguration `json:"configurations,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListSecurityDatasourceConfigurationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityDatasourceConfigurationsResponse struct{}"
	}

	return strings.Join([]string{"ListSecurityDatasourceConfigurationsResponse", string(data)}, " ")
}
