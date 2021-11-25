package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// checkType为cmdLine时checkMode填contain,checkContent格式为[“xxx”]表示进程命令行参数中需要包含xxx。checkType为env时checkMode填contain,checkContent格式为 [\"k1\",\"v1\"]表示进程环境变量中需要包含名为k1值为v1的环境变量。checkType为scope时checkMode填equals,checkContent格式为节点ID数组[\"hostId1”,”hostId2”],表示规则仅会在这些节点上生效(如果不指定节点范围,规则将下发到该项目所有的节点)。 规则发现部分,数组中有多个对象时表示需要同时满足所有条件的进程才会被匹配到。
type DiscoveryRule struct {
	// 匹配值。

	CheckContent []string `json:"checkContent"`
	// contain、equals 匹配条件。

	CheckMode string `json:"checkMode"`
	// cmdLine、env、scope 匹配类型。

	CheckType string `json:"checkType"`
}

func (o DiscoveryRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiscoveryRule struct{}"
	}

	return strings.Join([]string{"DiscoveryRule", string(data)}, " ")
}
