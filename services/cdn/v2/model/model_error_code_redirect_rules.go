package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ErrorCodeRedirectRules 自定义错误页面。
type ErrorCodeRedirectRules struct {

	// **参数解释：** 重定向的错误码 **约束限制：** 不涉及 **取值范围：** - 4xx: 400, 403, 404, 405, 414, 416, 451 - 5xx: 500, 501, 502, 503, 504  **默认取值：** 不涉及
	ErrorCode int32 `json:"error_code"`

	// **参数解释：** 执行规则 **约束限制：** 不涉及 **取值范围：** - break：如果错误码匹配了当前配置，请求将被重定向到目标Path。执行完当前规则后，当存在其他配置规则时，将不再匹配剩余规则。 - redirect：如果错误码匹配了当前配置，请求将被重定向到目标Path。执行完当前规则后，当存在其他配置规则时，将继续匹配剩余规则。  **默认取值：** 不涉及
	ExecutionMode *string `json:"execution_mode,omitempty"`

	// **参数解释：** 重定向状态码 **约束限制：** 当执行规则选择redirect时，需要配置该参数 **取值范围：** - 301 - 302  **默认取值：** 不涉及
	TargetCode int32 `json:"target_code"`

	// **参数解释：** 重定向的目标链接 **约束限制：** “执行规则”选择“Break”时：全路径匹配，支持输入一个目标地址，以“/”作为首字符，字符长度不超过512，如：/errorcode.html。 “执行规则”选择“Redirect”时：输入的URL须以http://或https:// 开头 ，字符长度不超过512，包含完整的域名和路径信息，如：http://example.com/errorcode.html。 **取值范围：** 不涉及 **默认取值：** 不涉及
	TargetLink string `json:"target_link"`
}

func (o ErrorCodeRedirectRules) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ErrorCodeRedirectRules struct{}"
	}

	return strings.Join([]string{"ErrorCodeRedirectRules", string(data)}, " ")
}
