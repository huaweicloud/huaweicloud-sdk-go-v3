package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowApplicableInstancesRequest Request Object
type ShowApplicableInstancesRequest struct {

	// 参数模板id
	ConfigId string `json:"config_id"`

	// 索引位置，偏移量。  从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询）。  取值必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询个数上限值。   - 取值范围: 1~100。   - 不传该参数时，默认查询前100条信息。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowApplicableInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowApplicableInstancesRequest struct{}"
	}

	return strings.Join([]string{"ShowApplicableInstancesRequest", string(data)}, " ")
}
