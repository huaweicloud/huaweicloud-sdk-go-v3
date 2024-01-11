package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSchemasRequest Request Object
type UpdateSchemasRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 数据库名称
	DatabaseName string `json:"database_name"`

	Body *WorkloadSchemaReq `json:"body,omitempty"`
}

func (o UpdateSchemasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSchemasRequest struct{}"
	}

	return strings.Join([]string{"UpdateSchemasRequest", string(data)}, " ")
}
