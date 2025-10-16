package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteAlertNoticeConfigRequest Request Object
type BatchDeleteAlertNoticeConfigRequest struct {

	// **参数解释：** 企业项目ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	EnterpriseProjectId string `json:"enterpriseProjectId"`

	// 语言，默认值为en-us。zh-cn（中文）/en-us（英文）
	XLanguage string `json:"X-Language"`

	Body *AlertNoticeConfigList `json:"body,omitempty"`
}

func (o BatchDeleteAlertNoticeConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteAlertNoticeConfigRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteAlertNoticeConfigRequest", string(data)}, " ")
}
