package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeWdrDto 节点WDR报表信息
type NodeWdrDto struct {

	// 节点WDR报表下载地址
	NodeWdr *string `json:"node_wdr,omitempty"`

	// 节点WDR报表名称
	NodeName *string `json:"node_name,omitempty"`
}

func (o NodeWdrDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeWdrDto struct{}"
	}

	return strings.Join([]string{"NodeWdrDto", string(data)}, " ")
}
