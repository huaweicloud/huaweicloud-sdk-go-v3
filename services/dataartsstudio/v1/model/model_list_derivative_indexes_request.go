package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListDerivativeIndexesRequest Request Object
type ListDerivativeIndexesRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 按名称或编码模糊查询
	Name *string `json:"name,omitempty"`

	// 按创建者查询
	CreateBy *string `json:"create_by,omitempty"`

	// 按审核人查询
	Approver *string `json:"approver,omitempty"`

	// 业务状态
	Status *ListDerivativeIndexesRequestStatus `json:"status,omitempty"`

	// 依据维度id查维度属性
	DimensionId *int64 `json:"dimension_id,omitempty"`

	// 依据维度颗粒度查维度属性
	DimensionGroup *string `json:"dimension_group,omitempty"`

	// 依据原子指标id查维度属性
	AtomicIndexId *int64 `json:"atomic_index_id,omitempty"`

	// 是否查询复合指标
	AllMetrics *bool `json:"all_metrics,omitempty"`

	// 数据连接类型
	DwType *string `json:"dw_type,omitempty"`

	// 业务对象l3 id
	L3Id *int64 `json:"l3_id,omitempty"`

	// 时间过滤左边界,与end_time一起使用,只支持时间范围过滤,单边过滤无效
	BeginTime *string `json:"begin_time,omitempty"`

	// 时间过滤右边界,与begin_time一起使用只支持时间范围过滤,单边过滤无效
	EndTime *string `json:"end_time,omitempty"`

	// 查询条数，即查询Y条数据。默认值50，取值范围[1,100]
	Limit *int32 `json:"limit,omitempty"`

	// 查询起始坐标，即跳过X条数据，仅支持0或limit的整数倍，不满足则向下取整。默认值0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListDerivativeIndexesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDerivativeIndexesRequest struct{}"
	}

	return strings.Join([]string{"ListDerivativeIndexesRequest", string(data)}, " ")
}

type ListDerivativeIndexesRequestStatus struct {
	value string
}

type ListDerivativeIndexesRequestStatusEnum struct {
	DRAFT              ListDerivativeIndexesRequestStatus
	PUBLISH_DEVELOPING ListDerivativeIndexesRequestStatus
	PUBLISHED          ListDerivativeIndexesRequestStatus
	OFFLINE_DEVELOPING ListDerivativeIndexesRequestStatus
	OFFLINE            ListDerivativeIndexesRequestStatus
	REJECT             ListDerivativeIndexesRequestStatus
}

func GetListDerivativeIndexesRequestStatusEnum() ListDerivativeIndexesRequestStatusEnum {
	return ListDerivativeIndexesRequestStatusEnum{
		DRAFT: ListDerivativeIndexesRequestStatus{
			value: "DRAFT",
		},
		PUBLISH_DEVELOPING: ListDerivativeIndexesRequestStatus{
			value: "PUBLISH_DEVELOPING",
		},
		PUBLISHED: ListDerivativeIndexesRequestStatus{
			value: "PUBLISHED",
		},
		OFFLINE_DEVELOPING: ListDerivativeIndexesRequestStatus{
			value: "OFFLINE_DEVELOPING",
		},
		OFFLINE: ListDerivativeIndexesRequestStatus{
			value: "OFFLINE",
		},
		REJECT: ListDerivativeIndexesRequestStatus{
			value: "REJECT",
		},
	}
}

func (c ListDerivativeIndexesRequestStatus) Value() string {
	return c.value
}

func (c ListDerivativeIndexesRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDerivativeIndexesRequestStatus) UnmarshalJSON(b []byte) error {
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
