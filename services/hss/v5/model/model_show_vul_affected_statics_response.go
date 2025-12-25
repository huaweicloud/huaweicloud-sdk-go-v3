package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVulAffectedStaticsResponse Response Object
type ShowVulAffectedStaticsResponse struct {

	// **参数解释**: 影响的漏洞数量(按漏洞编号计算) **取值范围**: 最小值0，最大值2147483647
	VulNum *int32 `json:"vul_num,omitempty"`

	// **参数解释**: 影响主机数量 **取值范围**: 最小值0，最大值2147483647
	HostNum *int32 `json:"host_num,omitempty"`

	// **参数解释**: 影响镜像数量 **取值范围**: 最小值0，最大值2147483647
	ImageNum *int32 `json:"image_num,omitempty"`

	// **参数解释**: 影响容器数量 **取值范围**: 最小值0，最大值2147483647
	ContainerNum *int32 `json:"container_num,omitempty"`

	// **参数解释**: 按漏洞类型的统计值，当select_type为all_host或空时，有该字段
	DataList *[]VulAffectedStatisticsResponseInfoDataList `json:"data_list,omitempty"`

	// **参数解释**: 影响的总漏洞条数(主机+漏洞 以此维度计算) **取值范围**: 最小值0，最大值2147483647
	TotalVulNum *int32 `json:"total_vul_num,omitempty"`

	// **参数解释**: 提示
	ExtendTips *[]string `json:"extend_tips,omitempty"`

	// **参数解释**: 漏洞修复提示 **取值范围**: 最小值1，最大值500
	ExtendTextTips *[]string `json:"extend_text_tips,omitempty"`

	DisabledOperateTypes *VulAffectedStatisticsResponseInfoDisabledOperateTypes `json:"disabled_operate_types,omitempty"`

	// **参数解释**: CCE漏洞数量 **取值范围**: 最小值0，最大值2147483647
	CceVulNum *int32 `json:"cce_vul_num,omitempty"`

	// **参数解释**: 基础版主机数量 **取值范围**: 最小值0，最大值2147483647
	BasicHostNum *int32 `json:"basic_host_num,omitempty"`

	// **参数解释**: CCE主机漏洞禁止修复列表
	CceDisabledVulList *[]VulAffectedStatisticsResponseInfoCceDisabledVulList `json:"cce_disabled_vul_list,omitempty"`
	HttpStatusCode     int                                                    `json:"-"`
}

func (o ShowVulAffectedStaticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVulAffectedStaticsResponse struct{}"
	}

	return strings.Join([]string{"ShowVulAffectedStaticsResponse", string(data)}, " ")
}
