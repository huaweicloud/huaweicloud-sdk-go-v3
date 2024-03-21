package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityDatasourceConfigurationsRequest Request Object
type ListSecurityDatasourceConfigurationsRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`
}

func (o ListSecurityDatasourceConfigurationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityDatasourceConfigurationsRequest struct{}"
	}

	return strings.Join([]string{"ListSecurityDatasourceConfigurationsRequest", string(data)}, " ")
}
