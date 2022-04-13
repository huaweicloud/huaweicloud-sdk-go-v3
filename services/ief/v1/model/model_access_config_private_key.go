package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OPC-UA协议下，采用证书认证时私钥字段，value字段为base64格式
type AccessConfigPrivateKey struct {
	// value 最大长度512， value允许英文字母、数字、下划线、中划线、点、逗号、@、#、+、\\、/、？、^、=、%、&、:、~

	Value string `json:"value"`
	// 标识属性是否可选，默认为true

	Optional *bool `json:"optional,omitempty"`

	Metadata *ValueInPropertyVisitorsRegisterTypeMetadata `json:"metadata,omitempty"`
}

func (o AccessConfigPrivateKey) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessConfigPrivateKey struct{}"
	}

	return strings.Join([]string{"AccessConfigPrivateKey", string(data)}, " ")
}
