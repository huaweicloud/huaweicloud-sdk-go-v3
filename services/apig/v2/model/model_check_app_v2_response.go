package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CheckAppV2Response struct {
	// 名称

	Name *string `json:"name,omitempty"`
	// 描述

	Remark *string `json:"remark,omitempty"`
	// 编号

	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CheckAppV2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CheckAppV2Response struct{}"
	}

	return strings.Join([]string{"CheckAppV2Response", string(data)}, " ")
}
