package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AgencyUsage struct {

	// 区域名称。
	Region string `json:"region"`

	// 统一资源名称列表。
	Resources []string `json:"resources"`
}

func (o AgencyUsage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgencyUsage struct{}"
	}

	return strings.Join([]string{"AgencyUsage", string(data)}, " ")
}
