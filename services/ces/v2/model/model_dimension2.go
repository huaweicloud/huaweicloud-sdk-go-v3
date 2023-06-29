package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Dimension2 指标维度
type Dimension2 struct {

	// 资源维度，如：弹性云服务器，则维度为instance_id；目前最大支持4个维度，各服务资源的指标维度名称可查看：“[服务指标维度](https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html)”。
	Name string `json:"name"`

	// 资源维度值，为资源的实例ID，如：4270ff17-aba3-4138-89fa-820594c39755。
	Value string `json:"value"`
}

func (o Dimension2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Dimension2 struct{}"
	}

	return strings.Join([]string{"Dimension2", string(data)}, " ")
}
