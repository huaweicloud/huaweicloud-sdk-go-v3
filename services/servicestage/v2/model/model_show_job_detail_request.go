package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowJobDetailRequest struct {

	// 部署任务ID。
	JobId string `json:"job_id" xml:"job_id"`

	// 应用组件实例ID。
	InstanceId *string `json:"instance_id,omitempty" xml:"instance_id"`

	// 指定查询的个数，可用于分页查询。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 指定查询的偏移量，可用于分页查询。
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 是否降序。true表示desc, false表示asc。
	Desc *string `json:"desc,omitempty" xml:"desc"`
}

func (o ShowJobDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowJobDetailRequest", string(data)}, " ")
}
