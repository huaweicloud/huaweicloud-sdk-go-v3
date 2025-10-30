package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VulAffectedStatisticsResponseInfoDataList struct {

	// **参数解释**: 漏洞类型 **取值范围**: - linux_vul：linux漏洞 - windows_vul：windows漏洞 - web_cms：Web-CMS漏洞 - app_vul：应用漏洞 - urgent_vul：应急漏洞
	Type *string `json:"type,omitempty"`

	// **参数解释**: 该漏洞类型的漏洞数量 **取值范围**: 最小值0，最大值2147483647
	VulNum *int32 `json:"vul_num,omitempty"`
}

func (o VulAffectedStatisticsResponseInfoDataList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulAffectedStatisticsResponseInfoDataList struct{}"
	}

	return strings.Join([]string{"VulAffectedStatisticsResponseInfoDataList", string(data)}, " ")
}
