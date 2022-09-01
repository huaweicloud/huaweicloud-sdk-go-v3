package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 创建状态回调请求体
type RecordRule struct {

	// 规则id，由服务端返回。创建或修改规则的时候不携带
	RuleId *string `json:"rule_id,omitempty" xml:"rule_id"`

	// 应用id
	AppId *string `json:"app_id,omitempty" xml:"app_id"`

	ObsAddr *RecordObsFileAddr `json:"obs_addr,omitempty" xml:"obs_addr"`

	// 录制格式：hls格式或者mp4格式
	RecordFormats *[]RecordRuleRecordFormats `json:"record_formats,omitempty" xml:"record_formats"`

	HlsConfig *HlsRecordConfig `json:"hls_config,omitempty" xml:"hls_config"`

	Mp4Config *Mp4RecordConfig `json:"mp4_config,omitempty" xml:"mp4_config"`

	// 创建时间，形如“2006-01-02T15:04:05.075Z”，时区为：UTC
	CreateTime *string `json:"create_time,omitempty" xml:"create_time"`

	// 更新时间，形如“2006-01-02T15:04:05.075Z”，时区为：UTC
	UpdateTime *string `json:"update_time,omitempty" xml:"update_time"`
}

func (o RecordRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecordRule struct{}"
	}

	return strings.Join([]string{"RecordRule", string(data)}, " ")
}

type RecordRuleRecordFormats struct {
	value string
}

type RecordRuleRecordFormatsEnum struct {
	HLS RecordRuleRecordFormats
	MP4 RecordRuleRecordFormats
}

func GetRecordRuleRecordFormatsEnum() RecordRuleRecordFormatsEnum {
	return RecordRuleRecordFormatsEnum{
		HLS: RecordRuleRecordFormats{
			value: "HLS",
		},
		MP4: RecordRuleRecordFormats{
			value: "MP4",
		},
	}
}

func (c RecordRuleRecordFormats) Value() string {
	return c.value
}

func (c RecordRuleRecordFormats) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RecordRuleRecordFormats) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
