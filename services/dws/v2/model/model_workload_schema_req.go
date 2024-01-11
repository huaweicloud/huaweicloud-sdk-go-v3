package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkloadSchemaReq 模式空间信息
type WorkloadSchemaReq struct {

	// 模式空间名称
	SchemaName string `json:"schema_name"`

	// Schema空间阈值
	PermSpace string `json:"perm_space"`
}

func (o WorkloadSchemaReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkloadSchemaReq struct{}"
	}

	return strings.Join([]string{"WorkloadSchemaReq", string(data)}, " ")
}
