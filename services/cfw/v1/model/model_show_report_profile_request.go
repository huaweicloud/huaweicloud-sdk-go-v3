package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowReportProfileRequest Request Object
type ShowReportProfileRequest struct {

	// **参数解释**： 防火墙ID，用户创建防火墙实例后产生的唯一ID，配置后可区分不同防火墙，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取 **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	FwInstanceId string `json:"fw_instance_id"`

	// **参数解释**： 安全报告模板ID，可以通过调用[查询安全报告模板列表接口]获得，通过返回值中的data.records.profile_id获得 **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	ReportProfileId string `json:"report_profile_id"`
}

func (o ShowReportProfileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReportProfileRequest struct{}"
	}

	return strings.Join([]string{"ShowReportProfileRequest", string(data)}, " ")
}
