package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListDashboardInfosRequest Request Object
type ListDashboardInfosRequest struct {

	// 企业项目Id
	EnterpriseId *string `json:"enterprise_id,omitempty"`

	// 指定企业项目下监控看板是否收藏，true:收藏，false:未收藏，填此参数时，enterprise_id必填
	IsFavorite *bool `json:"is_favorite,omitempty"`

	// 监控看板名称
	DashboardName *string `json:"dashboard_name,omitempty"`

	// 监控看板id
	DashboardId *string `json:"dashboard_id,omitempty"`

	// 监控看板类型, monitor_dashboard表示监控大盘,other表示自定义看板
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
