package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListHostRequest struct {
	// 页码

	Page *int32 `json:"page,omitempty"`
	// 单页条数

	Pagesize *int32 `json:"pagesize,omitempty"`
	// 域名

	Hostname *string `json:"hostname,omitempty"`
	// 策略名

	Policyname *string `json:"policyname,omitempty"`
}

func (o ListHostRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListHostRequest struct{}"
	}

	return strings.Join([]string{"ListHostRequest", string(data)}, " ")
}
