package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListEventRequest struct {
	// 查询日志的时间范围

	Recent ListEventRequestRecent `json:"recent"`
	// 域名id9从获取防护网站列表获取域名id）

	Hosts *[]string `json:"hosts,omitempty"`
	// 页码

	Page *int32 `json:"page,omitempty"`
	// 每页的条数

	Pagesize *int32 `json:"pagesize,omitempty"`
}

func (o ListEventRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListEventRequest struct{}"
	}

	return strings.Join([]string{"ListEventRequest", string(data)}, " ")
}

type ListEventRequestRecent struct {
	value string
}

type ListEventRequestRecentEnum struct {
	YESTERDAY ListEventRequestRecent
	TODAY     ListEventRequestRecent
	E_3DAYS   ListEventRequestRecent
	E_1WEEK   ListEventRequestRecent
	E_1MONTH  ListEventRequestRecent
}

func GetListEventRequestRecentEnum() ListEventRequestRecentEnum {
	return ListEventRequestRecentEnum{
		YESTERDAY: ListEventRequestRecent{
			value: "yesterday",
		},
		TODAY: ListEventRequestRecent{
			value: "today",
		},
		E_3DAYS: ListEventRequestRecent{
			value: "3days",
		},
		E_1WEEK: ListEventRequestRecent{
			value: "1week",
		},
		E_1MONTH: ListEventRequestRecent{
			value: "1month",
		},
	}
}

func (c ListEventRequestRecent) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListEventRequestRecent) UnmarshalJSON(b []byte) error {
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
