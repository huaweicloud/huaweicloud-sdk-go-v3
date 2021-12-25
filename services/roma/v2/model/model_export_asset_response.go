package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ExportAssetResponse struct {
	// 资产导出作业的ID，可用于查询作业进度，获取导出作业进度

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportAssetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportAssetResponse struct{}"
	}

	return strings.Join([]string{"ExportAssetResponse", string(data)}, " ")
}
