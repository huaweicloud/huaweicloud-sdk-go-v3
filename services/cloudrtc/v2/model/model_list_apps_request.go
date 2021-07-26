package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListAppsRequest struct {
	// 使用AK/SK方式认证时必选，携带的鉴权信息。

	Authorization *string `json:"Authorization,omitempty"`
	// 使用AK/SK方式认证时必选，请求的发生时间。

	XSdkDate *string `json:"X-Sdk-Date,omitempty"`
	// 使用AK/SK方式认证时必选，携带项目ID信息。

	XProjectId *string `json:"X-Project-Id,omitempty"`
	// 应用的状态：  - ACTIVATION：应用开启  - DEACTIVATION：应用停用  - ARREARS：应用欠费

	State *ListAppsRequestState `json:"state,omitempty"`
	// 查询结果起始编号，此处代表分页的页码，默认为0。

	Offset *int32 `json:"offset,omitempty"`
	// 查询结果集数量，此处代表每一页的条数，最小为1，最大为100。默认为100。

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAppsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListAppsRequest struct{}"
	}

	return strings.Join([]string{"ListAppsRequest", string(data)}, " ")
}

type ListAppsRequestState struct {
	value string
}

type ListAppsRequestStateEnum struct {
	ACTIVATION   ListAppsRequestState
	DEACTIVATION ListAppsRequestState
	ARREARS      ListAppsRequestState
}

func GetListAppsRequestStateEnum() ListAppsRequestStateEnum {
	return ListAppsRequestStateEnum{
		ACTIVATION: ListAppsRequestState{
			value: "ACTIVATION",
		},
		DEACTIVATION: ListAppsRequestState{
			value: "DEACTIVATION",
		},
		ARREARS: ListAppsRequestState{
			value: "ARREARS",
		},
	}
}

func (c ListAppsRequestState) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListAppsRequestState) UnmarshalJSON(b []byte) error {
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
