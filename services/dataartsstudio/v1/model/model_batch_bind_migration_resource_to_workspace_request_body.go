package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchBindMigrationResourceToWorkspaceRequestBody 批量关联或取消关联数据集成资源到工作空间请求体。
type BatchBindMigrationResourceToWorkspaceRequestBody struct {

	// 待关联或取消关联的资源列表。
	BandingResourceList []MigrationBindResource `json:"banding_resource_list"`
}

func (o BatchBindMigrationResourceToWorkspaceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchBindMigrationResourceToWorkspaceRequestBody struct{}"
	}

	return strings.Join([]string{"BatchBindMigrationResourceToWorkspaceRequestBody", string(data)}, " ")
}
