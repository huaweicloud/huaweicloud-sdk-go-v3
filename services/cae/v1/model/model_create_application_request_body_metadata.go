package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateApplicationRequestBodyMetadata 请求数据。
type CreateApplicationRequestBodyMetadata struct {

	// 应用名称。
	Name string `json:"name"`
}

func (o CreateApplicationRequestBodyMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApplicationRequestBodyMetadata struct{}"
	}

	return strings.Join([]string{"CreateApplicationRequestBodyMetadata", string(data)}, " ")
}
