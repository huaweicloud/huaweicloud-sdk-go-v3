package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAlertNoticeConfigRequest Request Object
type CreateAlertNoticeConfigRequest struct {

	// **参数解释：** 语言，默认值为en-us。zh-cn（中文）/en-us（英文） **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** zh-cn
	XLanguage string `json:"X-Language"`

	// **参数解释：** 企业项目ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	EnterpriseProjectId string `json:"enterpriseProjectId"`

	Body *AlertNoticeConfig `json:"body,omitempty"`
}

func (o CreateAlertNoticeConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAlertNoticeConfigRequest struct{}"
	}

	return strings.Join([]string{"CreateAlertNoticeConfigRequest", string(data)}, " ")
}
