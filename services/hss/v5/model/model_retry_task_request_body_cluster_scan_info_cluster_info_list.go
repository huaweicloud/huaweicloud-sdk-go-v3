package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RetryTaskRequestBodyClusterScanInfoClusterInfoList struct {

	// **参数解释**: 集群id **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释**： 该集群重新扫描的扫描项列表，若不指定则重新扫描集群下所有扫描失败的扫描项 **约束限制**: 不涉及 **取值范围**: 最小值1，最大值3
	ScanTypeList *[]string `json:"scan_type_list,omitempty"`
}

func (o RetryTaskRequestBodyClusterScanInfoClusterInfoList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryTaskRequestBodyClusterScanInfoClusterInfoList struct{}"
	}

	return strings.Join([]string{"RetryTaskRequestBodyClusterScanInfoClusterInfoList", string(data)}, " ")
}
