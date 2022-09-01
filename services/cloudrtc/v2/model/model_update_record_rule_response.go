package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type UpdateRecordRuleResponse struct {

	// 规则id，由服务端返回。创建或修改规则的时候不携带
	RuleId *string `json:"rule_id,omitempty" xml:"rule_id"`

	// 应用id
	AppId *string `json:"app_id,omitempty" xml:"app_id"`

	ObsAddr *RecordObsFileAddr `json:"obs_addr,omitempty" xml:"obs_addr"`

	// 录制格式：hls格式或者mp4格式
	RecordFormats *[]UpdateRecordRuleResponseRecordFormats `json:"record_formats,omitempty" xml:"record_formats"`

	HlsConfig *HlsRecordConfig `json:"hls_config,omitempty" xml:"hls_config"`

	Mp4Config *Mp4RecordConfig `json:"mp4_config,omitempty" xml:"mp4_config"`

	// 创建时间，形如“2006-01-02T15:04:05.075Z”，时区为：UTC
	CreateTime *string `json:"create_time,omitempty" xml:"create_time"`

	// 更新时间，形如“2006-01-02T15:04:05.075Z”，时区为：UTC
	UpdateTime *string `json:"update_time,omitempty" xml:"update_time"`

	XRequestId     *string `json:"X-request-Id,omitempty" xml:"X-request-Id"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateRecordRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRecordRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateRecordRuleResponse", string(data)}, " ")
}

type UpdateRecordRuleResponseRecordFormats struct {
	value string
}

type UpdateRecordRuleResponseRecordFormatsEnum struct {
	HLS UpdateRecordRuleResponseRecordFormats
	MP4 UpdateRecordRuleResponseRecordFormats
}

func GetUpdateRecordRuleResponseRecordFormatsEnum() UpdateRecordRuleResponseRecordFormatsEnum {
	return UpdateRecordRuleResponseRecordFormatsEnum{
		HLS: UpdateRecordRuleResponseRecordFormats{
			value: "HLS",
		},
		MP4: UpdateRecordRuleResponseRecordFormats{
			value: "MP4",
		},
	}
}

func (c UpdateRecordRuleResponseRecordFormats) Value() string {
	return c.value
}

func (c UpdateRecordRuleResponseRecordFormats) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateRecordRuleResponseRecordFormats) UnmarshalJSON(b []byte) error {
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
