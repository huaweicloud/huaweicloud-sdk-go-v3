package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OPC-UA协议下，采用证书认证时证书字段，value字段为base64格式
type AccessConfigCertificate struct {
	// value 最大长度512， value允许英文字母、数字、下划线、中划线、点、逗号、@、#、+、\\、/、？、^、=、%、&、:、~

	Value string `json:"value"`
	// 标识属性是否可选，默认为true

	Optional *bool `json:"optional,omitempty"`

	Metadata *ValueInPropertyVisitorsRegisterTypeMetadata `json:"metadata,omitempty"`
}

func (o AccessConfigCertificate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessConfigCertificate struct{}"
	}

	return strings.Join([]string{"AccessConfigCertificate", string(data)}, " ")
}
