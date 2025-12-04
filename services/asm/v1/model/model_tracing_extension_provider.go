package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TracingExtensionProvider struct {

	// provider实例name
	Name *string `json:"name,omitempty"`

	Zipkin *ZipkinTracingProvider `json:"zipkin,omitempty"`
}

func (o TracingExtensionProvider) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TracingExtensionProvider struct{}"
	}

	return strings.Join([]string{"TracingExtensionProvider", string(data)}, " ")
}
