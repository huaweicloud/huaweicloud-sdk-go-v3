package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteAssetResponse struct {
	// 资产导出或删除作业的ID，可用于查询作业进。

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteAssetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAssetResponse struct{}"
	}

	return strings.Join([]string{"DeleteAssetResponse", string(data)}, " ")
}
