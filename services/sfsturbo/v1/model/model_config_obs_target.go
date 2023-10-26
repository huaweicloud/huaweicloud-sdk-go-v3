package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfigObsTarget 配置 obs 后端信息
type ConfigObsTarget struct {

	// obs 配置名
	Name string `json:"name"`

	// obs 配置协议类型
	Type string `json:"type"`

	// obs 桶 endpoint
	Url string `json:"url"`

	// obs 桶名
	Bucket string `json:"bucket"`
}

func (o ConfigObsTarget) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigObsTarget struct{}"
	}

	return strings.Join([]string{"ConfigObsTarget", string(data)}, " ")
}
