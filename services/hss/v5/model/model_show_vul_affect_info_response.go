package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVulAffectInfoResponse Response Object
type ShowVulAffectInfoResponse struct {

	// **参数解释**: 漏洞影响的主机ID列表 **取值范围**: 最小值0，最大值65535
	AffectHostIds *[]string `json:"affect_host_ids,omitempty"`

	// **参数解释**: 主机下的漏洞ID列表 **取值范围**: 最小值0，最大值65535
	AffectVulIds   *[]string `json:"affect_vul_ids,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowVulAffectInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVulAffectInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowVulAffectInfoResponse", string(data)}, " ")
}
