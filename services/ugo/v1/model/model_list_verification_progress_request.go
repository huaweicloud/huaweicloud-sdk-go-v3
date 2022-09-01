package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListVerificationProgressRequest struct {

	// 迁移项目ID。
	MigrationProjectId string `json:"migration_project_id" xml:"migration_project_id"`
}

func (o ListVerificationProgressRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVerificationProgressRequest struct{}"
	}

	return strings.Join([]string{"ListVerificationProgressRequest", string(data)}, " ")
}
