package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListDimensionsRequest Request Object
type ListDimensionsRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 按名称或编码模糊查询
	Name *string `json:"name,omitempty"`

	// 按创建者查询
	CreateBy *string `json:"create_by,omitempty"`

	// 按审核人查询
	Approver *string `json:"approver,omitempty"`

	// 业务状态
	Status *ListDimensionsRequestStatus `json:"status,omitempty"`

	// 主题域l2 id
	L2Id *int64 `json:"l2_id,omitempty"`

	// 依据复合指标id列表查维度
	DerivativeIds *[]int64 `json:"derivative_ids,omitempty"`

	// 时间过滤左边界,与end_time一起使用,只支持时间范围过滤,单边过滤无效
	BeginTime *string `json:"begin_time,omitempty"`

	// 时间过滤右边界,与begin_time一起使用只支持时间范围过滤,单边过滤无效
	EndTime *string `json:"end_time,omitempty"`

	// 事实表id
	FactLogicId *int64 `json:"fact_logic_id,omitempty"`

	// 维度类型
	DimensionType *ListDimensionsRequestDimensionType `json:"dimension_type,omitempty"`

	// 查询条数，即查询Y条数据。默认值50，取值范围[1,100]
	Limit *int32 `json:"limit,omitempty"`

	// 查询起始坐标，即跳过X条数据，仅支持0或limit的整数倍，不满足则向下取整。默认值0
	Offset *int32 `json:"offset,omitempty"`

	// 所属的业务分层的id
	BizCatalogId *string `json:"biz_catalog_id,omitempty"`
}

func (o ListDimensionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDimensionsRequest struct{}"
	}

	return strings.Join([]string{"ListDimensionsRequest", string(data)}, " ")
}

type ListDimensionsRequestStatus struct {
	value string
}

type ListDimensionsRequestStatusEnum struct {
	DRAFT              ListDimensionsRequestStatus
	PUBLISH_DEVELOPING ListDimensionsRequestStatus
	PUBLISHED          ListDimensionsRequestStatus
	OFFLINE_DEVELOPING ListDimensionsRequestStatus
	OFFLINE            ListDimensionsRequestStatus
	REJECT             ListDimensionsRequestStatus
}

func GetListDimensionsRequestStatusEnum() ListDimensionsRequestStatusEnum {
	return ListDimensionsRequestStatusEnum{
		DRAFT: ListDimensionsRequestStatus{
			value: "DRAFT",
		},
		PUBLISH_DEVELOPING: ListDimensionsRequestStatus{
			value: "PUBLISH_DEVELOPING",
		},
		PUBLISHED: ListDimensionsRequestStatus{
			value: "PUBLISHED",
		},
		OFFLINE_DEVELOPING: ListDimensionsRequestStatus{
			value: "OFFLINE_DEVELOPING",
		},
		OFFLINE: ListDimensionsRequestStatus{
			value: "OFFLINE",
		},
		REJECT: ListDimensionsRequestStatus{
			value: "REJECT",
		},
	}
}

func (c ListDimensionsRequestStatus) Value() string {
	return c.value
}

func (c ListDimensionsRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDimensionsRequestStatus) UnmarshalJSON(b []byte) error {
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

type ListDimensionsRequestDimensionType struct {
	value string
}

type ListDimensionsRequestDimensionTypeEnum struct {
	COMMON      ListDimensionsRequestDimensionType
	LOOKUP      ListDimensionsRequestDimensionType
	HIERARCHIES ListDimensionsRequestDimensionType
}

func GetListDimensionsRequestDimensionTypeEnum() ListDimensionsRequestDimensionTypeEnum {
	return ListDimensionsRequestDimensionTypeEnum{
		COMMON: ListDimensionsRequestDimensionType{
			value: "COMMON",
		},
		LOOKUP: ListDimensionsRequestDimensionType{
			value: "LOOKUP",
		},
		HIERARCHIES: ListDimensionsRequestDimensionType{
			value: "HIERARCHIES",
		},
	}
}

func (c ListDimensionsRequestDimensionType) Value() string {
	return c.value
}

func (c ListDimensionsRequestDimensionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDimensionsRequestDimensionType) UnmarshalJSON(b []byte) error {
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
