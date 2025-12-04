package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHostProtectStatusRequestBody 修改域名防护状态请求体
type UpdateHostProtectStatusRequestBody struct {

	// **参数解释：** 域名防护状态标识，用于指定域名在WAF中的防护运行状态 **约束限制：** 不涉及 **取值范围：**  - -1：bypass，该域名的请求直接到达其后端服务器，不再经过WAF  - 0：暂停防护，WAF只转发该域名的请求，不做攻击检测  - 1：开启防护，WAF根据您配置的策略进行攻击检测  **默认取值：** 不涉及
	ProtectStatus int32 `json:"protect_status"`
}

func (o UpdateHostProtectStatusRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHostProtectStatusRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateHostProtectStatusRequestBody", string(data)}, " ")
}
