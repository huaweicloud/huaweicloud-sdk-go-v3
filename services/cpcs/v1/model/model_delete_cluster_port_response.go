package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteClusterPortResponse Response Object
type DeleteClusterPortResponse struct {

	// 删除资源结果
	DeleteTenantResource *string `json:"delete_tenant_resource,omitempty"`
	HttpStatusCode       int     `json:"-"`
}

func (o DeleteClusterPortResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteClusterPortResponse struct{}"
	}

	return strings.Join([]string{"DeleteClusterPortResponse", string(data)}, " ")
}
