package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateServerlessPolicy 设置serverless配置策略请求体。
type UpdateServerlessPolicy struct {

	// 单节点VCPUs伸缩下限，取值范围可根据[查询数据库规格](https://support.huaweicloud.com/api-gaussdb/ShowGaussMySqlFlavors.html)接口获取。
	MinVcpus int32 `json:"min_vcpus"`

	// 单节点VCPUs伸缩上限，取值范围可根据[查询数据库规格](https://support.huaweicloud.com/api-gaussdb/ShowGaussMySqlFlavors.html)接口获取。
	MaxVcpus int32 `json:"max_vcpus"`
}

func (o UpdateServerlessPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServerlessPolicy struct{}"
	}

	return strings.Join([]string{"UpdateServerlessPolicy", string(data)}, " ")
}
