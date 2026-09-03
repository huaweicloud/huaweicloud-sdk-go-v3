package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InvokeWdrReportResponse Response Object
type InvokeWdrReportResponse struct {

	// 实例WDR报表下载地址
	InstanceWdr *string `json:"instance_wdr,omitempty"`

	// WDR报表名称
	InstanceWdrName *string `json:"instance_wdr_name,omitempty"`

	// 节点WDR报表列表
	NodeWdrList *[]NodeWdrDto `json:"node_wdr_list,omitempty"`

	// WDR报表状态。取值范围：0（无报表）、1（生成中）、2（生成成功）、3（生成失败）
	WdrStatus *int32 `json:"wdr_status,omitempty"`

	// WDR快照开始ID
	StartSnapshotId *int64 `json:"start_snapshot_id,omitempty"`

	// WDR快照结束ID
	EndSnapshotId  *int64 `json:"end_snapshot_id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o InvokeWdrReportResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InvokeWdrReportResponse struct{}"
	}

	return strings.Join([]string{"InvokeWdrReportResponse", string(data)}, " ")
}
