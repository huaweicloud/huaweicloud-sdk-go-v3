package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListRenewRateOnPeriodReq struct {

	// |参数名称：资源ID列表。只支持传入主资源ID，最多10个资源ID。| |参数约束以及描述：资源ID列表。只支持传入主资源ID，最多10个资源ID。|
	ResourceIds []string `json:"resource_ids"`

	// |参数名称：周期类型：2：月3：年| |参数的约束及描述：周期类型：2：月3：年|
	PeriodType int32 `json:"period_type"`

	// |参数名称：周期数目：如果是月，目前支持1-11如果是年，目前支持1-3| |参数的约束及描述：周期数目：如果是月，目前支持1-11如果是年，目前支持1-3|
	PeriodNum int32 `json:"period_num"`

	// |参数名称：是否包含关联资源一起询价| |参数的约束及描述：该参数非必填，true:包含。false:不包含|
	IncludeRelativeResources *bool `json:"include_relative_resources,omitempty"`
}

func (o ListRenewRateOnPeriodReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRenewRateOnPeriodReq struct{}"
	}

	return strings.Join([]string{"ListRenewRateOnPeriodReq", string(data)}, " ")
}
