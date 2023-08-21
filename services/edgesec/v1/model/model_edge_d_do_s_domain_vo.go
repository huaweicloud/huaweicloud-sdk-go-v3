package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type EdgeDDoSDomainVo struct {

	// 域名id
	Id *string `json:"id,omitempty"`

	// 域名
	DomainName *string `json:"domain_name,omitempty"`

	// 租户ID
	TenantId *string `json:"tenant_id,omitempty"`

	// 域名在CDN所属区域
	AreaType *EdgeDDoSDomainVoAreaType `json:"area_type,omitempty"`

	// cdn域名调度情况（0：未防护，1：配置中，2：已防护，3：删除中）
	DispatchStatus *int32 `json:"dispatch_status,omitempty"`

	// 防护开关（0:关，1:开）
	ProtectedSwitch *int32 `json:"protected_switch,omitempty"`

	// 开启时间
	OpenDate *int64 `json:"open_date,omitempty"`

	// 关闭时间
	CloseDate *int64 `json:"close_date,omitempty"`

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o EdgeDDoSDomainVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EdgeDDoSDomainVo struct{}"
	}

	return strings.Join([]string{"EdgeDDoSDomainVo", string(data)}, " ")
}

type EdgeDDoSDomainVoAreaType struct {
	value string
}

type EdgeDDoSDomainVoAreaTypeEnum struct {
	MAINLAND_CHINA         EdgeDDoSDomainVoAreaType
	OUTSIDE_MAINLAND_CHINA EdgeDDoSDomainVoAreaType
	EUROPE                 EdgeDDoSDomainVoAreaType
	GLOBAL                 EdgeDDoSDomainVoAreaType
}

func GetEdgeDDoSDomainVoAreaTypeEnum() EdgeDDoSDomainVoAreaTypeEnum {
	return EdgeDDoSDomainVoAreaTypeEnum{
		MAINLAND_CHINA: EdgeDDoSDomainVoAreaType{
			value: "mainland_china",
		},
		OUTSIDE_MAINLAND_CHINA: EdgeDDoSDomainVoAreaType{
			value: "outside_mainland_china",
		},
		EUROPE: EdgeDDoSDomainVoAreaType{
			value: "europe",
		},
		GLOBAL: EdgeDDoSDomainVoAreaType{
			value: "global",
		},
	}
}

func (c EdgeDDoSDomainVoAreaType) Value() string {
	return c.value
}

func (c EdgeDDoSDomainVoAreaType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EdgeDDoSDomainVoAreaType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
