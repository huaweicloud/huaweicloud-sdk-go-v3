package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResponsesCreate struct {

	// 响应名称。支持英文、数字、下划线、中划线，1-64个字符。
	Name string `json:"name"`

	// 错误类型的响应定义，其中key为错误类型。key的枚举值为： - AUTH_FAILURE：认证失败 - AUTH_HEADER_MISSING：认证身份来源缺失 - AUTHORIZER_FAILURE：自定义认证失败 - AUTHORIZER_CONF_FAILURE：自定义认证配置错误 - AUTHORIZER_IDENTITIES_FAILURE：自定义认证身份来源错误 - BACKEND_UNAVAILABLE：后端不可用 - BACKEND_TIMEOUT：后端超时 - THROTTLED：调用次数超出阈值 - UNAUTHORIZED：应用未授权 - ACCESS_DENIED：拒绝访问 - NOT_FOUND：未找到匹配的API - REQUEST_PARAMETERS_FAILURE：请求参数错误 - DEFAULT_4XX：默认4XX - DEFAULT_5XX：默认5XX - THIRD_AUTH_FAILURE: 第三方认证失败 - THIRD_AUTH_IDENTITIES_FAILURE: 第三方认证身份来源错误 - THIRD_AUTH_CONF_FAILURE: 第三方认证配置错误 - ORCHESTRATION_PARAMETER_NOT_FOUND: 没有入参进行参数编排规则匹配，参数编排失败 - ORCHESTRATION_FAILURE: 有入参进行参数编排规则匹配，但是匹配不上编排规则，参数编排失败  每项错误类型均为一个JSON体
	Responses map[string]ResponseInfo `json:"responses,omitempty"`
}

func (o ResponsesCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResponsesCreate struct{}"
	}

	return strings.Join([]string{"ResponsesCreate", string(data)}, " ")
}
