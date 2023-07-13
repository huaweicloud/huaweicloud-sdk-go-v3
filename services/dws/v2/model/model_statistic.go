package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Statistic 资源数量详情。
type Statistic struct {

	// 资源名称。 - cluster.total：总集群（个）。 - cluster.normal：可用集群（个）。 - instance.total：总节点（个）。 - instance.normal：可用节点（个）。 - storage.total：总容量（GB）。
	Name string `json:"name"`

	// 资源数量值。
	Value float64 `json:"value"`

	// 资源数量单位。
	Unit string `json:"unit"`
}

func (o Statistic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Statistic struct{}"
	}

	return strings.Join([]string{"Statistic", string(data)}, " ")
}
