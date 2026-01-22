package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IpsWhitelistRespData struct {

	// **参数解释**：  白名单ID列表 **取值范围**：  不涉及
	ListIds *[]string `json:"list_ids,omitempty"`
}

func (o IpsWhitelistRespData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpsWhitelistRespData struct{}"
	}

	return strings.Join([]string{"IpsWhitelistRespData", string(data)}, " ")
}
