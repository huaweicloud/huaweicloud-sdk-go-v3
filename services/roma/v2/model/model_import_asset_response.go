package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportAssetResponse Response Object
type ImportAssetResponse struct {

	// 资产导入作业的ID，可用于查询作业进度，获取导入作业进度
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ImportAssetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportAssetResponse struct{}"
	}

	return strings.Join([]string{"ImportAssetResponse", string(data)}, " ")
}
