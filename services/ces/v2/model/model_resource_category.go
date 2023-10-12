package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceCategory struct {

	// 查询服务的命名空间，各服务命名空间请参考[服务命名空间](https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html)
	Namespace string `json:"namespace"`

	// 资源的维度信息，多个维度通过字母序排序后逗号拼接
	DimensionNames []string `json:"dimension_names"`
}

func (o ResourceCategory) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceCategory struct{}"
	}

	return strings.Join([]string{"ResourceCategory", string(data)}, " ")
}
