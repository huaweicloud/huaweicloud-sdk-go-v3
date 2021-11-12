package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 集群连接信息
type ClusterLinks struct {
	// 关系

	Rel *string `json:"rel,omitempty"`
	// 链接地址

	Href *string `json:"href,omitempty"`
}

func (o ClusterLinks) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterLinks struct{}"
	}

	return strings.Join([]string{"ClusterLinks", string(data)}, " ")
}
