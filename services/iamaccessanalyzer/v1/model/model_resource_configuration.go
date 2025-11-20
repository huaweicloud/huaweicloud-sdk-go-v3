package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceConfiguration 提权访问中单个资源的配置。
type ResourceConfiguration struct {

	// 资源的唯一资源标识符。
	Resource string `json:"resource"`

	// 当前资源要分析的操作列表。
	Actions []string `json:"actions"`
}

func (o ResourceConfiguration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceConfiguration struct{}"
	}

	return strings.Join([]string{"ResourceConfiguration", string(data)}, " ")
}
