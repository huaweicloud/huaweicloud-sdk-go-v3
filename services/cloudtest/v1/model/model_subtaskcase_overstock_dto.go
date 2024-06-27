package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SubtaskcaseOverstockDto struct {

	// 查询时间
	DataTime *sdktime.SdkTime `json:"data_time,omitempty"`

	// 执行机类型
	ExecutorType *string `json:"executor_type,omitempty"`

	// UUID
	Id *string `json:"id,omitempty"`

	// 执行机标签
	Label *string `json:"label,omitempty"`

	// 执行机所属区域
	LocationId *string `json:"location_id,omitempty"`

	// 积压数量
	SubtaskcaseOverstockCount *int32 `json:"subtaskcase_overstock_count,omitempty"`

	// 服务ID
	TestServiceId *string `json:"test_service_id,omitempty"`
}

func (o SubtaskcaseOverstockDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubtaskcaseOverstockDto struct{}"
	}

	return strings.Join([]string{"SubtaskcaseOverstockDto", string(data)}, " ")
}
