package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListBandwidthsRequest struct {
	// 取值为上一页数据的最后一条记录的id，为空时为查询第一页

	Marker *string `json:"marker,omitempty"`
	// 功能说明：每页返回的个数  取值范围：0~intmax

	Limit *int32 `json:"limit,omitempty"`
	// 功能说明：企业项目ID。可以使用该字段过滤某个企业项目下的虚拟私有云。  取值范围：最大长度36字节，带“-”连字符的UUID格式，或者是字符串“0”。“0”表示默认企业项目。若需要查询当前用户所有企业项目绑定的虚拟私有云，请传参all_granted_eps。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 功能说明：带宽类型，标识是否是共享带宽 取值范围：WHOLE，PER WHOLE表示共享带宽；PER，表示独享带宽

	ShareType *ListBandwidthsRequestShareType `json:"share_type,omitempty"`
}

func (o ListBandwidthsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListBandwidthsRequest struct{}"
	}

	return strings.Join([]string{"ListBandwidthsRequest", string(data)}, " ")
}

type ListBandwidthsRequestShareType struct {
	value string
}

type ListBandwidthsRequestShareTypeEnum struct {
	WHOLE ListBandwidthsRequestShareType
	PER   ListBandwidthsRequestShareType
}

func GetListBandwidthsRequestShareTypeEnum() ListBandwidthsRequestShareTypeEnum {
	return ListBandwidthsRequestShareTypeEnum{
		WHOLE: ListBandwidthsRequestShareType{
			value: "WHOLE",
		},
		PER: ListBandwidthsRequestShareType{
			value: "PER",
		},
	}
}

func (c ListBandwidthsRequestShareType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListBandwidthsRequestShareType) UnmarshalJSON(b []byte) error {
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
