package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Response Object
type ShowRecordRuleResponse struct {
	// 规则id，由服务端返回。创建或修改规则的时候不携带

	RuleId *string `json:"rule_id,omitempty"`
	// 应用id

	AppId *string `json:"app_id,omitempty"`

	ObsAddr *RecordObsFileAddr `json:"obs_addr,omitempty"`
	// 录制格式：hls格式或者mp4格式

	RecordFormats *[]ShowRecordRuleResponseRecordFormats `json:"record_formats,omitempty"`

	HlsConfig *HlsRecordConfig `json:"hls_config,omitempty"`

	Mp4Config *Mp4RecordConfig `json:"mp4_config,omitempty"`
	// 创建时间，形如“2006-01-02T15:04:05.075Z”，时区为：UTC

	CreateTime *string `json:"create_time,omitempty"`
	// 更新时间，形如“2006-01-02T15:04:05.075Z”，时区为：UTC

	UpdateTime *string `json:"update_time,omitempty"`

	XRequestId     *string `json:"X-request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowRecordRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRecordRuleResponse struct{}"
	}

	return strings.Join([]string{"ShowRecordRuleResponse", string(data)}, " ")
}

type ShowRecordRuleResponseRecordFormats struct {
	value string
}

type ShowRecordRuleResponseRecordFormatsEnum struct {
	HLS ShowRecordRuleResponseRecordFormats
	MP4 ShowRecordRuleResponseRecordFormats
}

func GetShowRecordRuleResponseRecordFormatsEnum() ShowRecordRuleResponseRecordFormatsEnum {
	return ShowRecordRuleResponseRecordFormatsEnum{
		HLS: ShowRecordRuleResponseRecordFormats{
			value: "HLS",
		},
		MP4: ShowRecordRuleResponseRecordFormats{
			value: "MP4",
		},
	}
}

func (c ShowRecordRuleResponseRecordFormats) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowRecordRuleResponseRecordFormats) UnmarshalJSON(b []byte) error {
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
