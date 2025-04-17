package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceDimension 指标维度
type ResourceDimension struct {

	// 资源维度，如：弹性云服务器，则维度为instance_id；目前最大支持4个维度，各服务资源的指标维度名称可查看：“[服务指标维度](ces_03_0059.xml)”。
	Name string `json:"name"`

	// 资源维度值，为资源的实例ID，如：4270ff17-aba3-4138-89fa-820594c39755。
	Value string `json:"value"`
}

func (o ResourceDimension) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceDimension struct{}"
	}

	return strings.Join([]string{"ResourceDimension", string(data)}, " ")
}
