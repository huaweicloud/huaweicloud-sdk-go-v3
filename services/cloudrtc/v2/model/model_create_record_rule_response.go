package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type CreateRecordRuleResponse struct {

	// 规则id，由服务端返回。创建或修改规则的时候不携带
	RuleId *string `json:"rule_id,omitempty" xml:"rule_id"`

	// 应用id
	AppId *string `json:"app_id,omitempty" xml:"app_id"`

	ObsAddr *RecordObsFileAddr `json:"obs_addr,omitempty" xml:"obs_addr"`

	// 录制格式：hls格式或者mp4格式
	RecordFormats *[]CreateRecordRuleResponseRecordFormats `json:"record_formats,omitempty" xml:"record_formats"`

	HlsConfig *HlsRecordConfig `json:"hls_config,omitempty" xml:"hls_config"`

	Mp4Config *Mp4RecordConfig `json:"mp4_config,omitempty" xml:"mp4_config"`

	// 创建时间，形如“2006-01-02T15:04:05.075Z”，时区为：UTC
	CreateTime *string `json:"create_time,omitempty" xml:"create_time"`

	// 更新时间，形如“2006-01-02T15:04:05.075Z”，时区为：UTC
	UpdateTime *string `json:"update_time,omitempty" xml:"update_time"`

	XRequestId     *string `json:"X-request-Id,omitempty" xml:"X-request-Id"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateRecordRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRecordRuleResponse struct{}"
	}

	return strings.Join([]string{"CreateRecordRuleResponse", string(data)}, " ")
}

type CreateRecordRuleResponseRecordFormats struct {
	value string
}

type CreateRecordRuleResponseRecordFormatsEnum struct {
	HLS CreateRecordRuleResponseRecordFormats
	MP4 CreateRecordRuleResponseRecordFormats
}

func GetCreateRecordRuleResponseRecordFormatsEnum() CreateRecordRuleResponseRecordFormatsEnum {
	return CreateRecordRuleResponseRecordFormatsEnum{
		HLS: CreateRecordRuleResponseRecordFormats{
			value: "HLS",
		},
		MP4: CreateRecordRuleResponseRecordFormats{
			value: "MP4",
		},
	}
}

func (c CreateRecordRuleResponseRecordFormats) Value() string {
	return c.value
}

func (c CreateRecordRuleResponseRecordFormats) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateRecordRuleResponseRecordFormats) UnmarshalJSON(b []byte) error {
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
