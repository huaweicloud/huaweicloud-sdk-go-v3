package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAlarmRulesRequest Request Object
type ListAlarmRulesRequest struct {

	// 告警规则ID
	AlarmId *string `json:"alarm_id,omitempty"`

	// 告警名称, 只能包含0-9/a-z/A-Z/_/-或汉字，长度1-128
	Name *string `json:"name,omitempty"`

	// 查询服务的命名空间，各服务命名空间请参考“[服务命名空间](ces_03_0059.xml)”
	Namespace *string `json:"namespace,omitempty"`

	// 告警资源ID，多维度情况按字母升序排列并使用逗号分隔
	ResourceId *string `json:"resource_id,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 产品层级跨纬规则查询时支持产品名称查询，一般由\"服务命名空间,服务首层维度名称\"组成，如\"SYS.ECS,instance_id\"
	ProductName *string `json:"product_name,omitempty"`

	// 产品层级跨纬规则查询时支持规则所属类型查询，resource_level取值为product即为产品层级跨纬规则，不填或者取值为dimension则为旧的规则类型
	ResourceLevel *ListAlarmRulesRequestResourceLevel `json:"resource_level,omitempty"`

	// 分页偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 分页大小
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAlarmRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmRulesRequest struct{}"
	}

	return strings.Join([]string{"ListAlarmRulesRequest", string(data)}, " ")
}

type ListAlarmRulesRequestResourceLevel struct {
	value string
}

type ListAlarmRulesRequestResourceLevelEnum struct {
	PRODUCT   ListAlarmRulesRequestResourceLevel
	DIMENSION ListAlarmRulesRequestResourceLevel
}

func GetListAlarmRulesRequestResourceLevelEnum() ListAlarmRulesRequestResourceLevelEnum {
	return ListAlarmRulesRequestResourceLevelEnum{
		PRODUCT: ListAlarmRulesRequestResourceLevel{
			value: "product",
		},
		DIMENSION: ListAlarmRulesRequestResourceLevel{
			value: "dimension",
		},
	}
}

func (c ListAlarmRulesRequestResourceLevel) Value() string {
	return c.value
}

func (c ListAlarmRulesRequestResourceLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAlarmRulesRequestResourceLevel) UnmarshalJSON(b []byte) error {
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
