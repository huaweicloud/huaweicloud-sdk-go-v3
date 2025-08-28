package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBotMTopnRequestResponse Response Object
type ListBotMTopnRequestResponse struct {

	// **参数解释：** TopN已知BOT的请求统计 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	TopnKnownBots *[]TypedStatBucket `json:"topn_known_bots,omitempty"`

	// **参数解释：** TopN被访问域名的请求统计 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	TopnDomains *[]TypedStatBucket `json:"topn_domains,omitempty"`

	// **参数解释：** TopN攻击源IP的请求统计 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	TopnSrcIps *[]TypedStatBucket `json:"topn_src_ips,omitempty"`

	// **参数解释：** TopN攻击源IP所属ASN的请求统计 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	TopnSrcIpAsns *[]TypedStatBucket `json:"topn_src_ip_asns,omitempty"`

	// **参数解释：** TopN攻击源地区的请求统计 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	TopnSrcIpGeolocations *[]TypedStatBucket `json:"topn_src_ip_geolocations,omitempty"`

	// **参数解释：** TopN JA4指纹的请求统计 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	TopnJa4Fps     *[]TypedStatBucket `json:"topn_ja4_fps,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListBotMTopnRequestResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBotMTopnRequestResponse struct{}"
	}

	return strings.Join([]string{"ListBotMTopnRequestResponse", string(data)}, " ")
}
