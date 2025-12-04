package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ZipkinTracingProvider struct {

	// zipkin服务地址
	Service *string `json:"service,omitempty"`

	// zipkin服务端口
	Port *int32 `json:"port,omitempty"`
}

func (o ZipkinTracingProvider) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ZipkinTracingProvider struct{}"
	}

	return strings.Join([]string{"ZipkinTracingProvider", string(data)}, " ")
}
