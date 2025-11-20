package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteVulWhiteListBody struct {

	// **参数解释**: 漏洞白名单列表 **约束限制**: 不涉及 **取值范围**: 最少1条，最多50条 **默认取值**: 不涉及
	Ids []string `json:"ids"`
}

func (o DeleteVulWhiteListBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVulWhiteListBody struct{}"
	}

	return strings.Join([]string{"DeleteVulWhiteListBody", string(data)}, " ")
}
