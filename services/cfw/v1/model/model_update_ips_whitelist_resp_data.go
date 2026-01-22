package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateIpsWhitelistRespData **参数解释** 返回数据 **取值范围**： 不涉及
type UpdateIpsWhitelistRespData struct {

	// **参数解释** 白名单id  **取值范围**： 不涉及
	IpsWhitelistId *string `json:"ips_whitelist_id,omitempty"`
}

func (o UpdateIpsWhitelistRespData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIpsWhitelistRespData struct{}"
	}

	return strings.Join([]string{"UpdateIpsWhitelistRespData", string(data)}, " ")
}
