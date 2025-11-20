package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAgencyResponse Response Object
type DeleteAgencyResponse struct {

	// **参数解释：** 代理id **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释：** 代理名称 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释：** 版本 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Version *string `json:"version,omitempty"`

	// **参数解释：** 代理存在时间段 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Duration *string `json:"duration,omitempty"`

	// **参数解释：** 使用代理的domainid **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	DomainId *string `json:"domain_id,omitempty"`

	// **参数解释：** 代理是否合法 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	IsValid        *bool `json:"is_valid,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o DeleteAgencyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAgencyResponse struct{}"
	}

	return strings.Join([]string{"DeleteAgencyResponse", string(data)}, " ")
}
