package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReportProfileRespData **参数解释**： 响应信息 **取值范围**： 不涉及
type ReportProfileRespData struct {

	// **参数解释**： 模板ID **取值范围**： 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**： 模板名称 **取值范围**： 不涉及
	Name *string `json:"name,omitempty"`
}

func (o ReportProfileRespData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReportProfileRespData struct{}"
	}

	return strings.Join([]string{"ReportProfileRespData", string(data)}, " ")
}
