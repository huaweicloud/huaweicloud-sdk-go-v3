package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDomainTemplateRequest Request Object
type DeleteDomainTemplateRequest struct {

	// **参数解释：** 域名模板ID，可以通过查询域名模板列表接口获取 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	TmlId string `json:"tml_id"`
}

func (o DeleteDomainTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDomainTemplateRequest struct{}"
	}

	return strings.Join([]string{"DeleteDomainTemplateRequest", string(data)}, " ")
}
