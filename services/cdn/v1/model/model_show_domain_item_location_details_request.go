package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ShowDomainItemLocationDetailsRequest struct {
	// 当用户开启企业项目功能时，该参数生效，表示资源所属企业项目，不传表示默认项目。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 查询开始时间戳，必须设为5分钟整时刻点

	StartTime int64 `json:"start_time"`
	// 查询结束时间戳，必须设为5分钟整时刻点，与开始时间戳时间差不可以超过一天

	EndTime int64 `json:"end_time"`
	// 域名列表，多个域名以逗号（半角）分隔，如：www.test1.com,www.test2.com，all表示查询名下全部域名

	DomainName string `json:"domain_name"`
	// 指标类型列表 网络资源消耗：bw（带宽），flux（流量） 访问情况：req_num（请求总数）

	StatType ShowDomainItemLocationDetailsRequestStatType `json:"stat_type"`
	// 区域列表，以逗号分隔，all表示查询全部区域

	Region string `json:"region"`
	// 运营商列表，以逗号分隔，all表示查询全部运营商

	Isp string `json:"isp"`
}

func (o ShowDomainItemLocationDetailsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDomainItemLocationDetailsRequest struct{}"
	}

	return strings.Join([]string{"ShowDomainItemLocationDetailsRequest", string(data)}, " ")
}

type ShowDomainItemLocationDetailsRequestStatType struct {
	value string
}

type ShowDomainItemLocationDetailsRequestStatTypeEnum struct {
	BW      ShowDomainItemLocationDetailsRequestStatType
	FLUX    ShowDomainItemLocationDetailsRequestStatType
	REQ_NUM ShowDomainItemLocationDetailsRequestStatType
}

func GetShowDomainItemLocationDetailsRequestStatTypeEnum() ShowDomainItemLocationDetailsRequestStatTypeEnum {
	return ShowDomainItemLocationDetailsRequestStatTypeEnum{
		BW: ShowDomainItemLocationDetailsRequestStatType{
			value: "bw",
		},
		FLUX: ShowDomainItemLocationDetailsRequestStatType{
			value: "flux",
		},
		REQ_NUM: ShowDomainItemLocationDetailsRequestStatType{
			value: "req_num",
		},
	}
}

func (c ShowDomainItemLocationDetailsRequestStatType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ShowDomainItemLocationDetailsRequestStatType) UnmarshalJSON(b []byte) error {
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
