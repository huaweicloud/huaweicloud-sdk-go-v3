package model

import (
	"encoding/json"

	"strings"
)

// 规则触发条件请求结构体
type AddRuleReq struct {
	// 用户自定义的规则名称。

	RuleName *string `json:"rule_name,omitempty"`
	// 用户自定义的规则描述。

	Description *string `json:"description,omitempty"`

	Subject *RoutingRuleSubject `json:"subject"`
	// 租户规则的生效范围，默认GLOBAL，取值如下： - GLOBAL：生效范围为租户级 - APP：生效范围为资源空间级。如果类型为APP，创建的规则生效范围为携带的app_id指定的资源空间，不携带app_id则创建规则生效范围为[默认资源空间](https://support.huaweicloud.com/usermanual-iothub/iot_01_0006.html#section0)。

	AppType *string `json:"app_type,omitempty"`
	// 资源空间ID。

	AppId *string `json:"app_id,omitempty"`
	// 用户自定义sql select语句，最大长度500，该参数仅供标准版和企业版用户使用。

	Select *string `json:"select,omitempty"`
	// 用户自定义sql where语句，最大长度500，该参数仅供标准版和企业版用户使用。

	Where *string `json:"where,omitempty"`
	// 规则推送消息的消息格式版本，租户无需设置，仅供内部兼容历史推送消息使用，iocsa的历史推送消息版本为V5.0.1

	DataVersion *string `json:"data_version,omitempty"`
}

func (o AddRuleReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AddRuleReq struct{}"
	}

	return strings.Join([]string{"AddRuleReq", string(data)}, " ")
}
