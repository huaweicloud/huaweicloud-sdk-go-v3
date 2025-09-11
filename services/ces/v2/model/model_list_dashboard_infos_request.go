package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListDashboardInfosRequest Request Object
type ListDashboardInfosRequest struct {

	// **参数解释**： 企业项目ID。 **约束限制**： 不涉及。 **取值范围**： 只能包含小写字母、数字、“-”、“_”，可以自定义企业项目ID，长度为36个字符。也可以为0（代表默认企业项目ID），all_granted_eps（代表所有企业项目ID）。           **默认取值**： 不涉及。
	EnterpriseId *string `json:"enterprise_id,omitempty"`

	// **参数解释**： 指定企业项目下监控看板是否收藏。 **约束限制**： 填此参数时，enterprise_id必填。 **取值范围**： - true:收藏 - false:未收藏          **默认取值**： 不涉及。
	IsFavorite *bool `json:"is_favorite,omitempty"`

	// **参数解释**： 监控看板名称。 **约束限制**： 不涉及。 **取值范围**： 长度为[1,128]个字符，只允许中文、英文、数字0-9、_和-          **默认取值**： 不涉及。
	DashboardName *string `json:"dashboard_name,omitempty"`

	// **参数解释**： 监控看板id。 **约束限制**： 不涉及。 **取值范围**： 以db开头，包含22个字母和数字，长度为24个字符 **默认取值**： 不涉及。
	DashboardId *string `json:"dashboard_id,omitempty"`

	// **参数解释**： 监控看板类型。 **约束限制**： 不涉及。 **取值范围**： - monitor_dashboard:表示监控大盘 - other:表示自定义看板 **默认取值**： 不涉及。
	DashboardType *ListDashboardInfosRequestDashboardType `json:"dashboard_type,omitempty"`
}

func (o ListDashboardInfosRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDashboardInfosRequest struct{}"
	}

	return strings.Join([]string{"ListDashboardInfosRequest", string(data)}, " ")
}

type ListDashboardInfosRequestDashboardType struct {
	value string
}

type ListDashboardInfosRequestDashboardTypeEnum struct {
	MONITOR_DASHBOARD ListDashboardInfosRequestDashboardType
	OTHER             ListDashboardInfosRequestDashboardType
}

func GetListDashboardInfosRequestDashboardTypeEnum() ListDashboardInfosRequestDashboardTypeEnum {
	return ListDashboardInfosRequestDashboardTypeEnum{
		MONITOR_DASHBOARD: ListDashboardInfosRequestDashboardType{
			value: "monitor_dashboard",
		},
		OTHER: ListDashboardInfosRequestDashboardType{
			value: "other",
		},
	}
}

func (c ListDashboardInfosRequestDashboardType) Value() string {
	return c.value
}

func (c ListDashboardInfosRequestDashboardType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDashboardInfosRequestDashboardType) UnmarshalJSON(b []byte) error {
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
