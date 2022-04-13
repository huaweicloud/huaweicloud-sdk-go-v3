package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 数据搜索引擎类型。
type ClusterListDatastore struct {
	// 支持类型：elasticsearch。

	Type *string `json:"type,omitempty"`
	// 引擎版本号。当前引擎版本为5.5.1、6.2.3、6.5.4、7.1.1、7.6.2、7.9.3。

	Version *string `json:"version,omitempty"`
}

func (o ClusterListDatastore) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterListDatastore struct{}"
	}

	return strings.Join([]string{"ClusterListDatastore", string(data)}, " ")
}
