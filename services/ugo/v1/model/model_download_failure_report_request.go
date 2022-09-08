package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DownloadFailureReportRequest struct {

	// 迁移项目ID。
	MigrationProjectId string `json:"migration_project_id"`
}

func (o DownloadFailureReportRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadFailureReportRequest struct{}"
	}

	return strings.Join([]string{"DownloadFailureReportRequest", string(data)}, " ")
}
