package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRestorableInstancesRequest Request Object
type ListRestorableInstancesRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 源实例id，需要恢复的实例ID。
	SourceInstanceId string `json:"source_instance_id"`

	// 实例备份信息ID，根据备份ID查询实例拓扑信息，过滤查询出来的实例，包含节点数，副本数等。参数为空时，根据restore_time查询。。
	BackupId *string `json:"backup_id,omitempty"`

	// 恢复点，当备份ID为空时，通过此参数查询实例拓扑信息，过滤实例列表。
	RestoreTime *string `json:"restore_time,omitempty"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询记录数。默认为100，不能为负数，最小值为1，最大值为100。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListRestorableInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRestorableInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListRestorableInstancesRequest", string(data)}, " ")
}
