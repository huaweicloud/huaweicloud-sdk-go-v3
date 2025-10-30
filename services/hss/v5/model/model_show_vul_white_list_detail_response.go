package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVulWhiteListDetailResponse Response Object
type ShowVulWhiteListDetailResponse struct {

	// **参数解释**: 白名单id **取值范围**: 字符长度0-256
	Id *string `json:"id,omitempty"`

	// **参数解释**: 漏洞id **取值范围**: 字符长度0-256
	VulId *string `json:"vul_id,omitempty"`

	// **参数解释**: 漏洞名称 **取值范围**: 字符长度0-256
	VulName *string `json:"vul_name,omitempty"`

	// **参数解释**: 漏洞类型 **取值范围**: - linux_vul：linux漏洞 - windows_vul：windows漏洞 - web_cms：Web-CMS漏洞 - app_vul：应用漏洞
	VulType *string `json:"vul_type,omitempty"`

	// **参数解释**: 漏洞对应的cve列表 **取值范围**: 最小值0，最大值2147483647
	Cves *[]VulWhiteListDetailResponseInfoCves `json:"cves,omitempty"`

	// **参数解释**: 白名单规则类型 **取值范围**: - all_host：白名单应用于全部主机 - specific_host：白名单应用于指定主机
	RuleType *string `json:"rule_type,omitempty"`

	// **参数解释**: 白名单规则为“指定主机”时，指定的主机列表 **取值范围**: 最小值0，最大值2147483647
	Hosts *[]VulWhiteListDetailResponseInfoHosts `json:"hosts,omitempty"`

	// **参数解释**: 白名单的描述信息 **取值范围**: 字符长度0-1000
	Description    *string `json:"description,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowVulWhiteListDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVulWhiteListDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowVulWhiteListDetailResponse", string(data)}, " ")
}
