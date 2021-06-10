package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ShowDomainItemDetailsRequest struct {
	// 当用户开启企业项目功能时，该参数生效，表示资源所属企业项目，不传表示默认项目。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 查询起始时间戳，必须设为5分钟整时刻点

	StartTime int64 `json:"start_time"`
	// 查询结束时间戳，必须设为5分钟整时刻点，与开始时间戳时间差不可以超过一天

	EndTime int64 `json:"end_time"`
	// 域名列表，多个域名以逗号（半角）分隔，如：www.test1.com,www.test2.com，all表示查询名下全部域名

	DomainName string `json:"domain_name"`
	// 网络资源消耗： - bw（带宽） - flux（流量） - bs_bw(回源带宽) - bs_flux（回源流量）  访问情况： - req_num（请求总数） - hit_num（请求命中次数） - bs_num(回源总数) - bs_fail_num(回源失败数) - hit_flux（命中流量）

	StatType ShowDomainItemDetailsRequestStatType `json:"stat_type"`
}

func (o ShowDomainItemDetailsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDomainItemDetailsRequest struct{}"
	}

	return strings.Join([]string{"ShowDomainItemDetailsRequest", string(data)}, " ")
}

type ShowDomainItemDetailsRequestStatType struct {
	value string
}

type ShowDomainItemDetailsRequestStatTypeEnum struct {
	BW          ShowDomainItemDetailsRequestStatType
	FLUX        ShowDomainItemDetailsRequestStatType
	BS_BW       ShowDomainItemDetailsRequestStatType
	BS_FLUX     ShowDomainItemDetailsRequestStatType
	REQ_NUM     ShowDomainItemDetailsRequestStatType
	HIT_NUM     ShowDomainItemDetailsRequestStatType
	BS_NUM      ShowDomainItemDetailsRequestStatType
	BS_FAIL_NUM ShowDomainItemDetailsRequestStatType
	HIT_FLUX    ShowDomainItemDetailsRequestStatType
}

func GetShowDomainItemDetailsRequestStatTypeEnum() ShowDomainItemDetailsRequestStatTypeEnum {
	return ShowDomainItemDetailsRequestStatTypeEnum{
		BW: ShowDomainItemDetailsRequestStatType{
			value: "bw",
		},
		FLUX: ShowDomainItemDetailsRequestStatType{
			value: "flux",
		},
		BS_BW: ShowDomainItemDetailsRequestStatType{
			value: "bs_bw",
		},
		BS_FLUX: ShowDomainItemDetailsRequestStatType{
			value: "bs_flux",
		},
		REQ_NUM: ShowDomainItemDetailsRequestStatType{
			value: "req_num",
		},
		HIT_NUM: ShowDomainItemDetailsRequestStatType{
			value: "hit_num",
		},
		BS_NUM: ShowDomainItemDetailsRequestStatType{
			value: "bs_num",
		},
		BS_FAIL_NUM: ShowDomainItemDetailsRequestStatType{
			value: "bs_fail_num",
		},
		HIT_FLUX: ShowDomainItemDetailsRequestStatType{
			value: "hit_flux",
		},
	}
}

func (c ShowDomainItemDetailsRequestStatType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ShowDomainItemDetailsRequestStatType) UnmarshalJSON(b []byte) error {
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
