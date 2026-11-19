package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStatisticsResponse Response Object
type ShowStatisticsResponse struct {

	// **参数解释**: 勒索病毒防护开启主机数量 **取值范围**: 最小值0，最大值2147483647
	VaultHostNum *int32 `json:"vault_host_num,omitempty"`

	// **参数解释**: 应用防护开启主机数量 **取值范围**: 最小值0，最大值2147483647
	RaspHostNum *int32 `json:"rasp_host_num,omitempty"`

	// **参数解释**: 网页防篡改开启主机数量 **取值范围**: 最小值0，最大值2147483647
	WtpHostNum *int32 `json:"wtp_host_num,omitempty"`

	// **参数解释**: 双因子认证开启主机数量 **取值范围**: 最小值0，最大值2147483647
	TwoFactorHostNum *int32 `json:"two_factor_host_num,omitempty"`

	// **参数解释**: 隔离文件数量 **取值范围**: 最小值0，最大值2147483647
	IsolatedFileNum *int32 `json:"isolated_file_num,omitempty"`

	// **参数解释**: 支持开启双因子的主机数量 **取值范围**: 最小值0，最大值2147483647
	SupportTwoFactorHostNum *int32 `json:"support_two_factor_host_num,omitempty"`

	// **参数解释**: 开启应用进程的主机数量 **取值范围**: 最小值0，最大值2147483647
	AppProcessHostNum *int32 `json:"app_process_host_num,omitempty"`
	HttpStatusCode    int    `json:"-"`
}

func (o ShowStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowStatisticsResponse", string(data)}, " ")
}
