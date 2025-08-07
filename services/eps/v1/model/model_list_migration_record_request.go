package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMigrationRecordRequest Request Object
type ListMigrationRecordRequest struct {

	// 开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 索引位置，从offset指定的下一条数据开始查询
	Offset *string `json:"offset,omitempty"`

	// 查询记录数
	Limit *int32 `json:"limit,omitempty"`

	// 资源ID
	ResourceId *string `json:"resource_id,omitempty"`
}

func (o ListMigrationRecordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMigrationRecordRequest struct{}"
	}

	return strings.Join([]string{"ListMigrationRecordRequest", string(data)}, " ")
}
