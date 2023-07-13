package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListDimensionLogicTablesRequest Request Object
type ListDimensionLogicTablesRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 按名称或编码模糊查询
	Name *string `json:"name,omitempty"`

	// 按创建者查询
	CreateBy *string `json:"create_by,omitempty"`

	// 按审核人查询
	Approver *string `json:"approver,omitempty"`

	// 按负责人查询
	Owner *string `json:"owner,omitempty"`

	// 业务状态
	Status *ListDimensionLogicTablesRequestStatus `json:"status,omitempty"`

	SyncStatus *ListDimensionLogicTablesRequestSyncStatus `json:"sync_status,omitempty"`

	SyncKey *[]ListDimensionLogicTablesRequestSyncKey `json:"sync_key,omitempty"`

	// 主题域l2 id
	L2Id *int64 `json:"l2_id,omitempty"`

	// 依据维度id查维度属性
	DimensionId *int64 `json:"dimension_id,omitempty"`

	// 时间过滤左边界,与end_time一起使用,只支持时间范围过滤,单边过滤无效
	BeginTime *string `json:"begin_time,omitempty"`

	// 时间过滤右边界,与begin_time一起使用只支持时间范围过滤,单边过滤无效
	EndTime *string `json:"end_time,omitempty"`

	// 查询条数，即查询Y条数据。默认值50，取值范围[1,100]
	Limit *int32 `json:"limit,omitempty"`

	// 查询起始坐标，即跳过X条数据，仅支持0或limit的整数倍，不满足则向下取整。默认值0
	Offset *int32 `json:"offset,omitempty"`

	// 所属的业务分层的id
	BizCatalogId *string `json:"biz_catalog_id,omitempty"`

	// 维度类型
	DimensionType *ListDimensionLogicTablesRequestDimensionType `json:"dimension_type,omitempty"`
}

func (o ListDimensionLogicTablesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDimensionLogicTablesRequest struct{}"
	}

	return strings.Join([]string{"ListDimensionLogicTablesRequest", string(data)}, " ")
}

type ListDimensionLogicTablesRequestStatus struct {
	value string
}

type ListDimensionLogicTablesRequestStatusEnum struct {
	DRAFT              ListDimensionLogicTablesRequestStatus
	PUBLISH_DEVELOPING ListDimensionLogicTablesRequestStatus
	PUBLISHED          ListDimensionLogicTablesRequestStatus
	OFFLINE_DEVELOPING ListDimensionLogicTablesRequestStatus
	OFFLINE            ListDimensionLogicTablesRequestStatus
	REJECT             ListDimensionLogicTablesRequestStatus
}

func GetListDimensionLogicTablesRequestStatusEnum() ListDimensionLogicTablesRequestStatusEnum {
	return ListDimensionLogicTablesRequestStatusEnum{
		DRAFT: ListDimensionLogicTablesRequestStatus{
			value: "DRAFT",
		},
		PUBLISH_DEVELOPING: ListDimensionLogicTablesRequestStatus{
			value: "PUBLISH_DEVELOPING",
		},
		PUBLISHED: ListDimensionLogicTablesRequestStatus{
			value: "PUBLISHED",
		},
		OFFLINE_DEVELOPING: ListDimensionLogicTablesRequestStatus{
			value: "OFFLINE_DEVELOPING",
		},
		OFFLINE: ListDimensionLogicTablesRequestStatus{
			value: "OFFLINE",
		},
		REJECT: ListDimensionLogicTablesRequestStatus{
			value: "REJECT",
		},
	}
}

func (c ListDimensionLogicTablesRequestStatus) Value() string {
	return c.value
}

func (c ListDimensionLogicTablesRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDimensionLogicTablesRequestStatus) UnmarshalJSON(b []byte) error {
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

type ListDimensionLogicTablesRequestSyncStatus struct {
	value string
}

type ListDimensionLogicTablesRequestSyncStatusEnum struct {
	RUNNING         ListDimensionLogicTablesRequestSyncStatus
	NO_NEED         ListDimensionLogicTablesRequestSyncStatus
	SUMMARY_SUCCESS ListDimensionLogicTablesRequestSyncStatus
	SUMMARY_FAILED  ListDimensionLogicTablesRequestSyncStatus
}

func GetListDimensionLogicTablesRequestSyncStatusEnum() ListDimensionLogicTablesRequestSyncStatusEnum {
	return ListDimensionLogicTablesRequestSyncStatusEnum{
		RUNNING: ListDimensionLogicTablesRequestSyncStatus{
			value: "RUNNING",
		},
		NO_NEED: ListDimensionLogicTablesRequestSyncStatus{
			value: "NO_NEED",
		},
		SUMMARY_SUCCESS: ListDimensionLogicTablesRequestSyncStatus{
			value: "SUMMARY_SUCCESS",
		},
		SUMMARY_FAILED: ListDimensionLogicTablesRequestSyncStatus{
			value: "SUMMARY_FAILED",
		},
	}
}

func (c ListDimensionLogicTablesRequestSyncStatus) Value() string {
	return c.value
}

func (c ListDimensionLogicTablesRequestSyncStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDimensionLogicTablesRequestSyncStatus) UnmarshalJSON(b []byte) error {
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

type ListDimensionLogicTablesRequestSyncKey struct {
	value string
}

type ListDimensionLogicTablesRequestSyncKeyEnum struct {
	BUSINESS_ASSET  ListDimensionLogicTablesRequestSyncKey
	DATA_QUALITY    ListDimensionLogicTablesRequestSyncKey
	TECHNICAL_ASSET ListDimensionLogicTablesRequestSyncKey
	META_DATA_LINK  ListDimensionLogicTablesRequestSyncKey
	PHYSICAL_TABLE  ListDimensionLogicTablesRequestSyncKey
	DLF_TASK        ListDimensionLogicTablesRequestSyncKey
	MATERIALIZATION ListDimensionLogicTablesRequestSyncKey
}

func GetListDimensionLogicTablesRequestSyncKeyEnum() ListDimensionLogicTablesRequestSyncKeyEnum {
	return ListDimensionLogicTablesRequestSyncKeyEnum{
		BUSINESS_ASSET: ListDimensionLogicTablesRequestSyncKey{
			value: "BUSINESS_ASSET",
		},
		DATA_QUALITY: ListDimensionLogicTablesRequestSyncKey{
			value: "DATA_QUALITY",
		},
		TECHNICAL_ASSET: ListDimensionLogicTablesRequestSyncKey{
			value: "TECHNICAL_ASSET",
		},
		META_DATA_LINK: ListDimensionLogicTablesRequestSyncKey{
			value: "META_DATA_LINK",
		},
		PHYSICAL_TABLE: ListDimensionLogicTablesRequestSyncKey{
			value: "PHYSICAL_TABLE",
		},
		DLF_TASK: ListDimensionLogicTablesRequestSyncKey{
			value: "DLF_TASK",
		},
		MATERIALIZATION: ListDimensionLogicTablesRequestSyncKey{
			value: "MATERIALIZATION",
		},
	}
}

func (c ListDimensionLogicTablesRequestSyncKey) Value() string {
	return c.value
}

func (c ListDimensionLogicTablesRequestSyncKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDimensionLogicTablesRequestSyncKey) UnmarshalJSON(b []byte) error {
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

type ListDimensionLogicTablesRequestDimensionType struct {
	value string
}

type ListDimensionLogicTablesRequestDimensionTypeEnum struct {
	COMMON      ListDimensionLogicTablesRequestDimensionType
	LOOKUP      ListDimensionLogicTablesRequestDimensionType
	HIERARCHIES ListDimensionLogicTablesRequestDimensionType
}

func GetListDimensionLogicTablesRequestDimensionTypeEnum() ListDimensionLogicTablesRequestDimensionTypeEnum {
	return ListDimensionLogicTablesRequestDimensionTypeEnum{
		COMMON: ListDimensionLogicTablesRequestDimensionType{
			value: "COMMON",
		},
		LOOKUP: ListDimensionLogicTablesRequestDimensionType{
			value: "LOOKUP",
		},
		HIERARCHIES: ListDimensionLogicTablesRequestDimensionType{
			value: "HIERARCHIES",
		},
	}
}

func (c ListDimensionLogicTablesRequestDimensionType) Value() string {
	return c.value
}

func (c ListDimensionLogicTablesRequestDimensionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDimensionLogicTablesRequestDimensionType) UnmarshalJSON(b []byte) error {
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
