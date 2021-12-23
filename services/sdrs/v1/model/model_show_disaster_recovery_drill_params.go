package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 查询容灾演练数据结构
type ShowDisasterRecoveryDrillParams struct {
	// 容灾演练的ID。

	Id string `json:"id"`
	// 容灾演练的名称。

	Name string `json:"name"`
	// 容灾演练的状态。

	Status string `json:"status"`
	// 演练虚拟私有云id。

	DrillVpcId string `json:"drill_vpc_id"`
	// 创建时间。默认格式为：\"yyyy-MM-dd HH:mm:ss.SSS\"，例：\"2019-04-01 12:00:00.000\"。

	CreatedAt string `json:"created_at"`
	// 更新时间。默认格式为：\"yyyy-MM-dd HH:mm:ss.SSS\"，例：\"2019-04-01 12:00:00.000\"。

	UpdatedAt string `json:"updated_at"`
	// 保护组的ID。

	ServerGroupId string `json:"server_group_id"`
	// 演练云服务器列表。

	DrillServers []DrillServerParams `json:"drill_servers"`
}

func (o ShowDisasterRecoveryDrillParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDisasterRecoveryDrillParams struct{}"
	}

	return strings.Join([]string{"ShowDisasterRecoveryDrillParams", string(data)}, " ")
}
