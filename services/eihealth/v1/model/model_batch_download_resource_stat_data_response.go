package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchDownloadResourceStatDataResponse struct {

	// 资源统计数据列表
	MetricStatistics *[]ResourceStatDataRsp `json:"metric_statistics,omitempty"`
	HttpStatusCode   int                    `json:"-"`
}

func (o BatchDownloadResourceStatDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDownloadResourceStatDataResponse struct{}"
	}

	return strings.Join([]string{"BatchDownloadResourceStatDataResponse", string(data)}, " ")
}
