package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowImageWhiteListDetailResponse Response Object
type ShowImageWhiteListDetailResponse struct {

	// **参数解释**： 白名单ID **取值范围**： 字符长度0-64位
	Id *string `json:"id,omitempty"`

	// **参数解释**： 漏洞ID（只在查询漏洞白名单时返回） **取值范围**： 字符长度0-256位
	VulId *string `json:"vul_id,omitempty"`

	// **参数解释**： 漏洞名称（只在查询漏洞白名单时返回） **取值范围**： 字符长度0-256位
	VulName *string `json:"vul_name,omitempty"`

	// **参数解释**： 漏洞类型（只在查询漏洞白名单时返回） **取值范围**: - linux_vul：linux漏洞。 - app_vul：应用漏洞。
	VulType *string `json:"vul_type,omitempty"`

	// **参数解释**: 漏洞对应的CVE列表（只在查询漏洞白名单时返回） **取值范围**: 最小值0，最大值1000
	Cves *[]ImageWhiteListDetailResponseInfoCves `json:"cves,omitempty"`

	// 白名单规则类型，包含如下：   -all_images : 白名单应用于全部镜像   -specific_image_types : 白名单应用于指定镜像类型(仅用于指定仓库镜像类型)   -specific_images : 白名单应用于指定镜像
	RuleType *string `json:"rule_type,omitempty"`

	QueryInfo *ImageWhiteListDetailResponseInfoQueryInfo `json:"query_info,omitempty"`

	// 白名单规则为“specific_images”时，指定的镜像列表。只在查询镜像白名单详情时返回数据。
	ImageInfo *[]ImageWhiteListDetailResponseInfoImageInfo `json:"image_info,omitempty"`

	// **参数解释**： 白名单的描述信息 **取值范围**： 字符长度0-1024位
	Description    *string `json:"description,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowImageWhiteListDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageWhiteListDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowImageWhiteListDetailResponse", string(data)}, " ")
}
