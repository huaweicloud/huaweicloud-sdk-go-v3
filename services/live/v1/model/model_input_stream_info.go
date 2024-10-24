package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// InputStreamInfo 频道入流信息
type InputStreamInfo struct {

	// 频道入流协议 - FLV_PULL - RTMP_PUSH - HLS_PULL - SRT_PULL - SRT_PUSH
	InputProtocol InputStreamInfoInputProtocol `json:"input_protocol"`

	// 频道主源流信息。入流协议为RTMP_PUSH和SRT_PUSH时，非必填项。其他情况下，均为必填项。
	Sources *[]SourcesInfo `json:"sources,omitempty"`

	// 备入流数组，非必填项。如果有备入流，则主备入流必须保证路数、codec和分辨率均一致。入流协议为RTMP_PUSH时，无需填写。
	SecondarySources *[]SecondarySourcesInfo `json:"secondary_sources,omitempty"`

	FailoverConditions *FailoverConditions `json:"failover_conditions,omitempty"`

	// 当入流协议为HLS_PULL时，需要配置的最大带宽。  用户提供的拉流URL中，针对不同码率的音视频，均会携带带宽参数“BANDWIDTH”。 - 如果这里配置最大带宽，媒体直播服务从URL拉流时，会选择小于最大带宽且码率最大的音视频流，推流到源站。 - 如果这里未配置最大带宽，媒体直播服务从URL拉流时，会默认选择“BANDWIDTH”最高的音视频流，推流到源站。
	MaxBandwidthLimit *int32 `json:"max_bandwidth_limit,omitempty"`

	// 当推流协议为SRT_PUSH时，如果配置了直推源站，编码器不支持输入streamid，需要打开设置为true
	IpPortMode *bool `json:"ip_port_mode,omitempty"`

	// SRT_PUSH类型时，客户push ip白名单
	IpWhitelist *string `json:"ip_whitelist,omitempty"`

	// 广告的scte35信号源。  仅HLS_PULL类型的频道支持此配置，且目前仅支持SEGMENTS。
	Scte35Source *string `json:"scte35_source,omitempty"`

	// 广告触发器配置。  包含如下取值： - Splice insert：拼接插入 - Provider advertisement：提供商广告 - Distributor advertisement：分销商广告 - Provider placement opportunity：提供商置放机会 - Distributor placement opportunity：分销商置放机会
	AdTriggers *[]string `json:"ad_triggers,omitempty"`

	// 设置音频选择器，最多设置8个音频选择器
	AudioSelectors *[]InputAudioSelector `json:"audio_selectors,omitempty"`
}

func (o InputStreamInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InputStreamInfo struct{}"
	}

	return strings.Join([]string{"InputStreamInfo", string(data)}, " ")
}

type InputStreamInfoInputProtocol struct {
	value string
}

type InputStreamInfoInputProtocolEnum struct {
	FLV_PULL  InputStreamInfoInputProtocol
	RTMP_PUSH InputStreamInfoInputProtocol
	HLS_PULL  InputStreamInfoInputProtocol
	SRT_PULL  InputStreamInfoInputProtocol
	SRT_PUSH  InputStreamInfoInputProtocol
}

func GetInputStreamInfoInputProtocolEnum() InputStreamInfoInputProtocolEnum {
	return InputStreamInfoInputProtocolEnum{
		FLV_PULL: InputStreamInfoInputProtocol{
			value: "FLV_PULL",
		},
		RTMP_PUSH: InputStreamInfoInputProtocol{
			value: "RTMP_PUSH",
		},
		HLS_PULL: InputStreamInfoInputProtocol{
			value: "HLS_PULL",
		},
		SRT_PULL: InputStreamInfoInputProtocol{
			value: "SRT_PULL",
		},
		SRT_PUSH: InputStreamInfoInputProtocol{
			value: "SRT_PUSH",
		},
	}
}

func (c InputStreamInfoInputProtocol) Value() string {
	return c.value
}

func (c InputStreamInfoInputProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InputStreamInfoInputProtocol) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
