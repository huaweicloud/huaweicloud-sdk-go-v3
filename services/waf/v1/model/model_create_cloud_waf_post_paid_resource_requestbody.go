package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCloudWafPostPaidResourceRequestbody 开通云模式按需请求体
type CreateCloudWafPostPaidResourceRequestbody struct {

	// 租户所在的站点 - hec-hk：华为云国际站 - hws：华为云大陆站
	ConsoleArea string `json:"console_area"`

	// **参数解释：** 按需功能名称 **取值范围：**  - CLOUD_WAF：按需云模式  - LARGE_MODEL_FIREWALL_AI_GUARD_DETECT: AI安全护栏
	PostpaidName *string `json:"postpaid_name,omitempty"`

	// **参数解释：** 扩展参数 **取值范围：** 不涉及
	ExtendParams *interface{} `json:"extend_params,omitempty"`
}

func (o CreateCloudWafPostPaidResourceRequestbody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCloudWafPostPaidResourceRequestbody struct{}"
	}

	return strings.Join([]string{"CreateCloudWafPostPaidResourceRequestbody", string(data)}, " ")
}
