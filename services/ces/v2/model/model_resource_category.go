package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceCategory struct {

	// 服务的命名空间，查询各服务命名空间请参考“[服务命名空间](ces_03_0059.xml)”
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
