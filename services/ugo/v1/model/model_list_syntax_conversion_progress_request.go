package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSyntaxConversionProgressRequest Request Object
type ListSyntaxConversionProgressRequest struct {

	// 迁移项目ID。
	MigrationProjectId string `json:"migration_project_id"`
}

func (o ListSyntaxConversionProgressRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSyntaxConversionProgressRequest struct{}"
	}

	return strings.Join([]string{"ListSyntaxConversionProgressRequest", string(data)}, " ")
}
