package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CustomMetrics struct {
	Exec *Exec `json:"exec,omitempty"`

	HttpGet *HttpGet `json:"http_get,omitempty"`
}

func (o CustomMetrics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomMetrics struct{}"
	}

	return strings.Join([]string{"CustomMetrics", string(data)}, " ")
}
