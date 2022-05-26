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
	RuleId *string `json:"rule_id,omitempty"`

	// 应用id
	AppId *string `json:"app_id,omitempty"`

	ObsAddr *RecordObsFileAddr `json:"obs_addr,omitempty"`

	// 录制格式：hls格式或者mp4格式
	RecordFormats *[]RecordRuleRecordFormats `json:"record_formats,omitempty"`

	HlsConfig *HlsRecordConfig `json:"hls_config,omitempty"`

	Mp4Config *Mp4RecordConfig `json:"mp4_config,omitempty"`

	// 创建时间，形如“2006-01-02T15:04:05.075Z”，时区为：UTC
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间，形如“2006-01-02T15:04:05.075Z”，时区为：UTC
	UpdateTime *string `json:"update_time,omitempty"`
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
