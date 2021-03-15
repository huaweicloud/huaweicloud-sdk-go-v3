package model

import (
	"encoding/json"

	"strings"
)

type ModifyProxyWeightRequest struct {
	// 主实例权重
	MasterWeight string `json:"master_weight"`
	// 只读实例信息
	ReadonlyInstances []ProxyReadonlyInstances `json:"readonly_instances"`
}

func (o ModifyProxyWeightRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ModifyProxyWeightRequest struct{}"
	}

	return strings.Join([]string{"ModifyProxyWeightRequest", string(data)}, " ")
}
