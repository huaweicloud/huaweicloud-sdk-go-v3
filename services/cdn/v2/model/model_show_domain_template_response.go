package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDomainTemplateResponse Response Object
type ShowDomainTemplateResponse struct {

	// **参数解释：** 查询域名模板总数 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Total *int32 `json:"total,omitempty"`

	Temlates       *[]TemplateItem `json:"temlates,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowDomainTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainTemplateResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainTemplateResponse", string(data)}, " ")
}
