package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type MigrateResourceRequest struct {

	// 目标企业项目ID，enterprise_project_id为0时表示迁移资源到默认资源组default。
	EnterpriseProjectId string `json:"enterprise_project_id"`

	Body *MigrateResource `json:"body,omitempty"`
}

func (o MigrateResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateResourceRequest struct{}"
	}

	return strings.Join([]string{"MigrateResourceRequest", string(data)}, " ")
}
