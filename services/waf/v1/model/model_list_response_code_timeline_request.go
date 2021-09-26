package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListResponseCodeTimelineRequest struct {
	// 企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 起始时间

	From int64 `json:"from"`
	// 结束时间

	To int64 `json:"to"`
	// 要查询域名列表（通过ListHost接口查询）

	Hosts *string `json:"hosts,omitempty"`
	// 要查询实例列表

	Instances *string `json:"instances,omitempty"`
	// 响应源

	ResponseSource *ListResponseCodeTimelineRequestResponseSource `json:"response_source,omitempty"`
	// 展示维度，按天展示时传\"DAY\"

	GroupBy *string `json:"group_by,omitempty"`
}

func (o ListResponseCodeTimelineRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListResponseCodeTimelineRequest struct{}"
	}

	return strings.Join([]string{"ListResponseCodeTimelineRequest", string(data)}, " ")
}

type ListResponseCodeTimelineRequestResponseSource struct {
	value string
}

type ListResponseCodeTimelineRequestResponseSourceEnum struct {
	WAF      ListResponseCodeTimelineRequestResponseSource
	UPSTREAM ListResponseCodeTimelineRequestResponseSource
}

func GetListResponseCodeTimelineRequestResponseSourceEnum() ListResponseCodeTimelineRequestResponseSourceEnum {
	return ListResponseCodeTimelineRequestResponseSourceEnum{
		WAF: ListResponseCodeTimelineRequestResponseSource{
			value: "WAF",
		},
		UPSTREAM: ListResponseCodeTimelineRequestResponseSource{
			value: "UPSTREAM",
		},
	}
}

func (c ListResponseCodeTimelineRequestResponseSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListResponseCodeTimelineRequestResponseSource) UnmarshalJSON(b []byte) error {
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
