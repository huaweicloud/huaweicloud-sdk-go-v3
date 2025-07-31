package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTaskResourcesRequestBodyClusterScanInfo 集群扫描任务特有的查询条件，任务类型为“cluster_scan”可用
type ListTaskResourcesRequestBodyClusterScanInfo struct {

	// 集群的扫描状态，包含如下：   - scanning：扫描中   - success：扫描成功   - failed：扫描失败
	ScanStatus *string `json:"scan_status,omitempty"`

	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`

	// 集群id
	ClusterId *string `json:"cluster_id,omitempty"`

	// 集群类型，包含如下： - cce：CCE集群 - ali：阿里云集群 - tencent：腾讯云集群 - azure：微软云集群 - aws：亚马逊集群 - self_built_hw：华为云自建集群 - self_built_idc：IDC自建集群
	ClusterType *string `json:"cluster_type,omitempty"`
}

func (o ListTaskResourcesRequestBodyClusterScanInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTaskResourcesRequestBodyClusterScanInfo struct{}"
	}

	return strings.Join([]string{"ListTaskResourcesRequestBodyClusterScanInfo", string(data)}, " ")
}
