package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSupportBindingConnectionBandwidthsRequest Request Object
type ListSupportBindingConnectionBandwidthsRequest struct {

	// 每页返回的个数。 取值范围：1~1000。
	Limit *int32 `json:"limit,omitempty"`

	// 翻页信息，从上次API调用返回的翻页数据中获取，可填写前一页marker或者后一页marker，填入前一页previous_marker就向前翻页，后一页next_marker就向翻页。 翻页过程中，查询条件不能修改，包括过滤条件，排序条件，limit。
	Marker *string `json:"marker,omitempty"`

	// 根据企业项目ID过滤列表。
	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`

	// 功能说明：本端接入点。   如果是region类型，则返回所有满足条件的城域带宽，不进行该字段的匹配过滤   如果是其他类型，则会用该字段跟全域互联带宽的local_area进行匹配过滤   附带过滤条件：会通过local_area和remote_area推算最佳全域互联带宽类型进行过滤查询   限制：local_area和remote_area同为空或者同不为空，且站点类型需一致
	LocalArea *string `json:"local_area,omitempty"`

	// 功能说明：远端接入点。   如果是region类型，则返回所有满足条件的城域带宽，不进行该字段的匹配过滤   如果是其他类型，则会用该字段跟全域互联带宽的remote_area进行匹配过滤   附带过滤条件：会通过local_area和remote_area推算最佳全域互联带宽类型进行过滤查询   限制：local_area和remote_area同为空或者同不为空，且站点类型需一致
	RemoteArea *string `json:"remote_area,omitempty"`

	// 根据支持绑定实例类型过滤全域互联带宽列表。实例类型： - CC: 云连接 - GEIP: 全域弹性公网IP - GCN: 中心网络 - GSN: 分支网络
	BindingService *ListSupportBindingConnectionBandwidthsRequestBindingService `json:"binding_service,omitempty"`
}

func (o ListSupportBindingConnectionBandwidthsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSupportBindingConnectionBandwidthsRequest struct{}"
	}

	return strings.Join([]string{"ListSupportBindingConnectionBandwidthsRequest", string(data)}, " ")
}

type ListSupportBindingConnectionBandwidthsRequestBindingService struct {
	value string
}

type ListSupportBindingConnectionBandwidthsRequestBindingServiceEnum struct {
	CC   ListSupportBindingConnectionBandwidthsRequestBindingService
	GEIP ListSupportBindingConnectionBandwidthsRequestBindingService
	GCN  ListSupportBindingConnectionBandwidthsRequestBindingService
	GSN  ListSupportBindingConnectionBandwidthsRequestBindingService
}

func GetListSupportBindingConnectionBandwidthsRequestBindingServiceEnum() ListSupportBindingConnectionBandwidthsRequestBindingServiceEnum {
	return ListSupportBindingConnectionBandwidthsRequestBindingServiceEnum{
		CC: ListSupportBindingConnectionBandwidthsRequestBindingService{
			value: "CC",
		},
		GEIP: ListSupportBindingConnectionBandwidthsRequestBindingService{
			value: "GEIP",
		},
		GCN: ListSupportBindingConnectionBandwidthsRequestBindingService{
			value: "GCN",
		},
		GSN: ListSupportBindingConnectionBandwidthsRequestBindingService{
			value: "GSN",
		},
	}
}

func (c ListSupportBindingConnectionBandwidthsRequestBindingService) Value() string {
	return c.value
}

func (c ListSupportBindingConnectionBandwidthsRequestBindingService) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSupportBindingConnectionBandwidthsRequestBindingService) UnmarshalJSON(b []byte) error {
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
