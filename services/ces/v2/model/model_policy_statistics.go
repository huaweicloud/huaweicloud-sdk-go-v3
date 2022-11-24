package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PolicyStatistics struct {

	// 查询服务的命名空间，各服务命名空间请参考“[服务命名空间](ces_03_0059.xml)”
	Namespace string `json:"namespace"`

	// 对应命名空间的告警策略数目
	PolicyNum int32 `json:"policy_num"`
}

func (o PolicyStatistics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyStatistics struct{}"
	}

	return strings.Join([]string{"PolicyStatistics", string(data)}, " ")
}
