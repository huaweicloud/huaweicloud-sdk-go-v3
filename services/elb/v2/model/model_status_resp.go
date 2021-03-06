package model

import (
	"encoding/json"

	"strings"
)

// 负载均衡器状态树对象
type StatusResp struct {
	Loadbalancer *LoadbalancerInStatusResp `json:"loadbalancer"`
}

func (o StatusResp) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "StatusResp struct{}"
	}

	return strings.Join([]string{"StatusResp", string(data)}, " ")
}
