package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CommitVerificationRequest Request Object
type CommitVerificationRequest struct {

	// 迁移项目ID。
	MigrationProjectId string `json:"migration_project_id"`
}

func (o CommitVerificationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommitVerificationRequest struct{}"
	}

	return strings.Join([]string{"CommitVerificationRequest", string(data)}, " ")
}
