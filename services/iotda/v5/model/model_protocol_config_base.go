package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProtocolConfigBase 泛协议配置信息
type ProtocolConfigBase struct {

	// **参数说明**：泛协议配置ID。
	ProtocolId *string `json:"protocol_id,omitempty"`

	// **参数说明**：泛协议配置的端口号。
	Port *int32 `json:"port,omitempty"`

	// **参数说明**：连接空闲断链时间，单位（s）。
	KeepaliveTimeout *int32 `json:"keepalive_timeout,omitempty"`

	// **参数说明**：协议类型。 **取值范围**： - TCP：通用TCP协议接入
	ProtocolType *string `json:"protocol_type,omitempty"`

	// **参数说明**：泛协议的描述信息。 **取值范围**：长度不超过2048，只允许中文、字母、数字、以及_?'#().,&%@!-等字符的组合
	Description *string `json:"description,omitempty"`

	// **参数说明**：是否开启tls加密传输。
	SslEnable *bool `json:"ssl_enable,omitempty"`

	// **参数说明**：码流拆包组包配置列表。
	FrameDecodeConfigs *[]FrameDecodeConfig `json:"frame_decode_configs,omitempty"`

	// **参数说明**：编解码类型。 **取值范围**： - FGS：将编解码插件以函数形式部署到FunctionGraph。 - PLUGIN：将编解码插件以OSGI插件形式部署到设备接入平台，使用该方式需提工单联系技术支持。
	CodecMode *string `json:"codec_mode,omitempty"`

	// **参数说明**：函数的URN（Uniform Resource Name），唯一标识函数，采用FGS进行编解码的对应函数地址。 **取值范围**：长度不超过256，只允许字母、数字、下划线（_）、连接符（-）、分隔符（:）的组合。
	FuncUrn *string `json:"func_urn,omitempty"`

	// **参数说明**：在物联网平台查询自定义鉴权的时间。格式：yyyyMMdd'T'HHmmss'Z'，如：20151212T121212Z。
	CreateTime *string `json:"create_time,omitempty"`

	// **参数说明**：在物联网平台更新查询自定义鉴权的时间。格式：yyyyMMdd'T'HHmmss'Z'，如：20151212T121212Z。
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o ProtocolConfigBase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtocolConfigBase struct{}"
	}

	return strings.Join([]string{"ProtocolConfigBase", string(data)}, " ")
}
