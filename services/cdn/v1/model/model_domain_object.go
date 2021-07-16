package model

import (
	"encoding/json"

	"strings"
)

type DomainObject struct {
	// 域名

	DomainName *string `json:"domain_name,omitempty"`
	// 数据结束时间戳，可能与请求时间不一致，可能不返回

	Flux *[]int64 `json:"flux,omitempty"`
}

func (o DomainObject) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DomainObject struct{}"
	}

	return strings.Join([]string{"DomainObject", string(data)}, " ")
}
