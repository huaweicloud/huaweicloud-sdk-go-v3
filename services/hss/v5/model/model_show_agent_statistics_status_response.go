package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAgentStatisticsStatusResponse Response Object
type ShowAgentStatisticsStatusResponse struct {

	// 主机Agent在线数
	Online *int32 `json:"online,omitempty"`

	// 主机Agent离线数
	Offline *int32 `json:"offline,omitempty"`

	// 未安装主机Agent的主机数
	NotInstalled *int32 `json:"not_installed,omitempty"`

	// 主机总数
	TotalNum       *int32 `json:"total_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowAgentStatisticsStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAgentStatisticsStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowAgentStatisticsStatusResponse", string(data)}, " ")
}
