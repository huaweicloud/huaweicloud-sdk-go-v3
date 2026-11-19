package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScansInfoRes 扫描信息
type ScansInfoRes struct {

	// **参数解释**： 漏洞统计列表 **取值范围**： 0-200个VulnerabilitiesStatisticsInfo对象
	VulnerabilitiesStatisticsInfoList *[]interface{} `json:"vulnerabilities_statistics_info_list,omitempty"`

	// **参数解释**： 资产账号变动记录列表 **取值范围**： 0-200个AccountChangeHistoryInfo对象
	AccountChangeHistoryInfoList *[]interface{} `json:"account_change_history_info_list,omitempty"`

	// **参数解释**： 危险开放端口列表 **取值范围**： 0-200个DangerPortsInfo对象
	DangerPortsInfoList *[]interface{} `json:"danger_ports_info_list,omitempty"`

	// **参数解释**： 弱口令列表 **取值范围**： 0-200个WeakPwdInfo对象
	WeakPwdInfoList *[]interface{} `json:"weak_pwd_info_list,omitempty"`
}

func (o ScansInfoRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScansInfoRes struct{}"
	}

	return strings.Join([]string{"ScansInfoRes", string(data)}, " ")
}
