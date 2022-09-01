package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowDetailsOfGatewayResponseV2Response struct {

	// 响应名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 错误类型的响应定义，其中key为错误类型。key的枚举值为： - AUTH_FAILURE：认证失败 - AUTH_HEADER_MISSING：认证身份来源缺失 - AUTHORIZER_FAILURE：自定义认证失败 - AUTHORIZER_CONF_FAILURE：自定义认证配置错误 - AUTHORIZER_IDENTITIES_FAILURE：自定义认证身份来源错误 - BACKEND_UNAVAILABLE：后端不可用 - BACKEND_TIMEOUT：后端超时 - THROTTLED：调用次数超出阈值 - UNAUTHORIZED：应用未授权 - ACCESS_DENIED：拒绝访问 - NOT_FOUND：未找到匹配的API - REQUEST_PARAMETERS_FAILURE：请求参数错误 - DEFAULT_4XX：默认4XX - DEFAULT_5XX：默认5XX  每项错误类型均为一个JSON体
	Responses map[string]ResponseInfoResp `json:"responses,omitempty" xml:"responses"`

	// 响应ID
	Id *string `json:"id,omitempty" xml:"id"`

	// 是否为分组默认响应
	Default *bool `json:"default,omitempty" xml:"default"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty" xml:"create_time"`

	// 修改时间
	UpdateTime     *sdktime.SdkTime `json:"update_time,omitempty" xml:"update_time"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowDetailsOfGatewayResponseV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailsOfGatewayResponseV2Response struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfGatewayResponseV2Response", string(data)}, " ")
}
