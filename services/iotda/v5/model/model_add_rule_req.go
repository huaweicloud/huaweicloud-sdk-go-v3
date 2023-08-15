package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddRuleReq 规则触发条件请求结构体
type AddRuleReq struct {

	// **参数说明**：规则名称。 **取值范围**：长度不超过256，只允许中文、字母、数字、以及_?'#().,&%@!-等字符的组合
	RuleName *string `json:"rule_name,omitempty"`

	// **参数说明**：用户自定义的规则描述。
	Description *string `json:"description,omitempty"`

	Subject *RoutingRuleSubject `json:"subject"`

	// **参数说明**：租户规则的生效范围，默认GLOBAL，。 **取值范围**： - GLOBAL：生效范围为租户级。 - APP：生效范围为资源空间级。如果类型为APP，创建的规则生效范围为携带的app_id指定的资源空间，不携带app_id则创建规则生效范围为[[默认资源空间](https://support.huaweicloud.com/usermanual-iothub/iot_01_0006.html#section0)](tag:hws)[[默认资源空间](https://support.huaweicloud.com/intl/zh-cn/usermanual-iothub/iot_01_0006.html#section0)](tag:hws_hk)。
	AppType *string `json:"app_type,omitempty"`

	// **参数说明**：资源空间ID。。 **取值范围**：长度不超过36，只允许字母、数字、下划线（_）、连接符（-）的组合。
	AppId *string `json:"app_id,omitempty"`

	// **参数说明**：用户自定义sql select语句，最大长度2500，该参数仅供标准版和企业版用户使用。可填参数可参考帮助文档数据下各接口的请求参数，如notify_data.body。
	Select *string `json:"select,omitempty"`

	// **参数说明**：用户自定义sql where语句，最大长度2500，该参数仅供标准版和企业版用户使用可填参数可参考帮助文档数据下各接口的请求参数，如notify_data.body。
	Where *string `json:"where,omitempty"`
}

func (o AddRuleReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddRuleReq struct{}"
	}

	return strings.Join([]string{"AddRuleReq", string(data)}, " ")
}
