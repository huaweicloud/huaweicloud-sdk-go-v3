package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MultiCloudClusterSyncStatusRequestBody 多云集群的同步接入状态请求
type MultiCloudClusterSyncStatusRequestBody struct {

	// 集群总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 多云集群的集群id列表
	DataList *[]MultiCloudClusterIdInfo `json:"data_list,omitempty"`
}

func (o MultiCloudClusterSyncStatusRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultiCloudClusterSyncStatusRequestBody struct{}"
	}

	return strings.Join([]string{"MultiCloudClusterSyncStatusRequestBody", string(data)}, " ")
}
