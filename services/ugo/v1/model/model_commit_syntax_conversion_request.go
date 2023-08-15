package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CommitSyntaxConversionRequest Request Object
type CommitSyntaxConversionRequest struct {

	// 迁移项目ID。
	MigrationProjectId string `json:"migration_project_id"`
}

func (o CommitSyntaxConversionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommitSyntaxConversionRequest struct{}"
	}

	return strings.Join([]string{"CommitSyntaxConversionRequest", string(data)}, " ")
}
