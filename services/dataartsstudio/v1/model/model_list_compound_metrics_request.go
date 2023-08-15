package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListCompoundMetricsRequest Request Object
type ListCompoundMetricsRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 按名称或编码模糊查询
	Name *string `json:"name,omitempty"`

	// 按创建者查询
	CreateBy *string `json:"create_by,omitempty"`

	// 按审核人查询
	Approver *string `json:"approver,omitempty"`

	// 业务状态
	Status *ListCompoundMetricsRequestStatus `json:"status,omitempty"`

	// 依据维度颗粒度查维度属性
	DimensionGroup *string `json:"dimension_group,omitempty"`

	// 依据原子指标id查维度属性
	AtomicIndexId *int64 `json:"atomic_index_id,omitempty"`

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

func (o ListCompoundMetricsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCompoundMetricsRequest struct{}"
	}

	return strings.Join([]string{"ListCompoundMetricsRequest", string(data)}, " ")
}

type ListCompoundMetricsRequestStatus struct {
	value string
}

type ListCompoundMetricsRequestStatusEnum struct {
	DRAFT              ListCompoundMetricsRequestStatus
	PUBLISH_DEVELOPING ListCompoundMetricsRequestStatus
	PUBLISHED          ListCompoundMetricsRequestStatus
	OFFLINE_DEVELOPING ListCompoundMetricsRequestStatus
	OFFLINE            ListCompoundMetricsRequestStatus
	REJECT             ListCompoundMetricsRequestStatus
}

func GetListCompoundMetricsRequestStatusEnum() ListCompoundMetricsRequestStatusEnum {
	return ListCompoundMetricsRequestStatusEnum{
		DRAFT: ListCompoundMetricsRequestStatus{
			value: "DRAFT",
		},
		PUBLISH_DEVELOPING: ListCompoundMetricsRequestStatus{
			value: "PUBLISH_DEVELOPING",
		},
		PUBLISHED: ListCompoundMetricsRequestStatus{
			value: "PUBLISHED",
		},
		OFFLINE_DEVELOPING: ListCompoundMetricsRequestStatus{
			value: "OFFLINE_DEVELOPING",
		},
		OFFLINE: ListCompoundMetricsRequestStatus{
			value: "OFFLINE",
		},
		REJECT: ListCompoundMetricsRequestStatus{
			value: "REJECT",
		},
	}
}

func (c ListCompoundMetricsRequestStatus) Value() string {
	return c.value
}

func (c ListCompoundMetricsRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListCompoundMetricsRequestStatus) UnmarshalJSON(b []byte) error {
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
