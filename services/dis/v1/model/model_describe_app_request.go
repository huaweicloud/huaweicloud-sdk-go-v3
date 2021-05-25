package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DescribeAppRequest struct {
	// 需要查询的App名称。

	AppName string `json:"app_name"`
}

func (o DescribeAppRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DescribeAppRequest struct{}"
	}

	return strings.Join([]string{"DescribeAppRequest", string(data)}, " ")
}
