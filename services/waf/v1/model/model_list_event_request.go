package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListEventRequest struct {
	// 您可以通过调用企业项目管理服务（EPS)的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 查询日志的时间范围,(不能和from、to同时使用)

	Recent *ListEventRequestRecent `json:"recent,omitempty"`
	// 起始时间(13位时间戳)，需要和to同时使用，不能和recent参数同时使用

	From *int64 `json:"from,omitempty"`
	// 结束时间(13位时间戳)，需要和from同时使用，不能和recent参数同时使用

	To *int64 `json:"to,omitempty"`
	// 域名id，从获取防护网站列表（ListHost）接口获取域名id

	Hosts *[]string `json:"hosts,omitempty"`
	// 分页查询时，返回第几页数据。范围0-100000，默认值为1，表示返回第1页数据。

	Page *int32 `json:"page,omitempty"`
	// 分页查询时，每页包含多少条结果。范围1-100，默认值为10，表示每页包含10条结果。

	Pagesize *int32 `json:"pagesize,omitempty"`
}

func (o ListEventRequest) String() string {
	data, err := utils.Marshal(o)
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
	return utils.Marshal(c.value)
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
