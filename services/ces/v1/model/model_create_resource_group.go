package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 资源分组中选择的资源信息。
type CreateResourceGroup struct {

	// 资源类型。即命名空间，如弹性云服务器的资源命名空间为：SYS.ECS；各服务命名空间可查看：“[服务命名空间](https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html)”。
	Namespace string `json:"namespace" xml:"namespace"`

	// 一个或者多个资源维度。
	Dimensions []MetricsDimension `json:"dimensions" xml:"dimensions"`
}

func (o CreateResourceGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourceGroup struct{}"
	}

	return strings.Join([]string{"CreateResourceGroup", string(data)}, " ")
}
