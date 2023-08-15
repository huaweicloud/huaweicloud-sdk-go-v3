package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Endpoint 端点详情
type Endpoint struct {
	Endpoint *EndpointObj `json:"endpoint,omitempty"`
}

func (o Endpoint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Endpoint struct{}"
	}

	return strings.Join([]string{"Endpoint", string(data)}, " ")
}
