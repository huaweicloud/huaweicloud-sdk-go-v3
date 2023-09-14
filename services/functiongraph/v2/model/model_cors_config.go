package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CorsConfig struct {

	// 是否支持跨域
	IsCors *bool `json:"is_cors,omitempty"`
}

func (o CorsConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CorsConfig struct{}"
	}

	return strings.Join([]string{"CorsConfig", string(data)}, " ")
}
