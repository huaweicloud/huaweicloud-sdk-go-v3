package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MetricsDimensionResp 指标维度
type MetricsDimensionResp struct {

	// **参数解释** 资源维度，如：弹性云服务器，则维度为instance_id；目前最大支持4个维度，各服务资源的指标维度名称可查看：“[服务指标维度](ces_03_0059.xml)”。 **取值范围** 由字母开头，后面可以包含字母、数字、_或-，长度为[1,32]个字符
	Name *string `json:"name,omitempty"`

	// **参数解释** 资源维度值，为资源的实例ID，如：4270ff17-aba3-4138-89fa-820594c39755。 **取值范围** 长度为[1,256]个字符
	Value *string `json:"value,omitempty"`
}

func (o MetricsDimensionResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricsDimensionResp struct{}"
	}

	return strings.Join([]string{"MetricsDimensionResp", string(data)}, " ")
}
