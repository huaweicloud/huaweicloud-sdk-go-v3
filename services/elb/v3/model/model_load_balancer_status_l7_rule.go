package model

import (
	"encoding/json"

	"strings"
)

// 查询负载均衡状态树返回对象中的rule数据模型
type LoadBalancerStatusL7Rule struct {
}

func (o LoadBalancerStatusL7Rule) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "LoadBalancerStatusL7Rule struct{}"
	}

	return strings.Join([]string{"LoadBalancerStatusL7Rule", string(data)}, " ")
}
