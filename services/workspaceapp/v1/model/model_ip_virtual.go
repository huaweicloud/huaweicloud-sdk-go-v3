package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IpVirtual 虚拟IP功能选项
type IpVirtual struct {

	// 是否启用虚拟IP功能。 启用虚拟IP功能将占用额外的ip地址，注意合理规划网络ip数量。 约束： * 只支持windows镜像。 * 只支持在创建服务器组时设置功能开关，不支持更新功能开关。 * 只支持具备dhcp动态分配ip能力的网络。
	Enable bool `json:"enable"`
}

func (o IpVirtual) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpVirtual struct{}"
	}

	return strings.Join([]string{"IpVirtual", string(data)}, " ")
}
