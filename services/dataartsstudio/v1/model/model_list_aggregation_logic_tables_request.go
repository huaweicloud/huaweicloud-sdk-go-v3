package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAggregationLogicTablesRequest Request Object
type ListAggregationLogicTablesRequest struct {

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
	Status *ListAggregationLogicTablesRequestStatus `json:"status,omitempty"`

	SyncStatus *ListAggregationLogicTablesRequestSyncStatus `json:"sync_status,omitempty"`

	SyncKey *[]ListAggregationLogicTablesRequestSyncKey `json:"sync_key,omitempty"`

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

	// 所属的业务分层的id
	BizCatalogId *string `json:"biz_catalog_id,omitempty"`

	// 是否自动生成
	AutoGenerate *bool `json:"auto_generate,omitempty"`
}

func (o ListAggregationLogicTablesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAggregationLogicTablesRequest struct{}"
	}

	return strings.Join([]string{"ListAggregationLogicTablesRequest", string(data)}, " ")
}

type ListAggregationLogicTablesRequestStatus struct {
	value string
}

type ListAggregationLogicTablesRequestStatusEnum struct {
	DRAFT              ListAggregationLogicTablesRequestStatus
	PUBLISH_DEVELOPING ListAggregationLogicTablesRequestStatus
	PUBLISHED          ListAggregationLogicTablesRequestStatus
	OFFLINE_DEVELOPING ListAggregationLogicTablesRequestStatus
	OFFLINE            ListAggregationLogicTablesRequestStatus
	REJECT             ListAggregationLogicTablesRequestStatus
}

func GetListAggregationLogicTablesRequestStatusEnum() ListAggregationLogicTablesRequestStatusEnum {
	return ListAggregationLogicTablesRequestStatusEnum{
		DRAFT: ListAggregationLogicTablesRequestStatus{
			value: "DRAFT",
		},
		PUBLISH_DEVELOPING: ListAggregationLogicTablesRequestStatus{
			value: "PUBLISH_DEVELOPING",
		},
		PUBLISHED: ListAggregationLogicTablesRequestStatus{
			value: "PUBLISHED",
		},
		OFFLINE_DEVELOPING: ListAggregationLogicTablesRequestStatus{
			value: "OFFLINE_DEVELOPING",
		},
		OFFLINE: ListAggregationLogicTablesRequestStatus{
			value: "OFFLINE",
		},
		REJECT: ListAggregationLogicTablesRequestStatus{
			value: "REJECT",
		},
	}
}

func (c ListAggregationLogicTablesRequestStatus) Value() string {
	return c.value
}

func (c ListAggregationLogicTablesRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAggregationLogicTablesRequestStatus) UnmarshalJSON(b []byte) error {
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

type ListAggregationLogicTablesRequestSyncStatus struct {
	value string
}

type ListAggregationLogicTablesRequestSyncStatusEnum struct {
	RUNNING         ListAggregationLogicTablesRequestSyncStatus
	NO_NEED         ListAggregationLogicTablesRequestSyncStatus
	SUMMARY_SUCCESS ListAggregationLogicTablesRequestSyncStatus
	SUMMARY_FAILED  ListAggregationLogicTablesRequestSyncStatus
}

func GetListAggregationLogicTablesRequestSyncStatusEnum() ListAggregationLogicTablesRequestSyncStatusEnum {
	return ListAggregationLogicTablesRequestSyncStatusEnum{
		RUNNING: ListAggregationLogicTablesRequestSyncStatus{
			value: "RUNNING",
		},
		NO_NEED: ListAggregationLogicTablesRequestSyncStatus{
			value: "NO_NEED",
		},
		SUMMARY_SUCCESS: ListAggregationLogicTablesRequestSyncStatus{
			value: "SUMMARY_SUCCESS",
		},
		SUMMARY_FAILED: ListAggregationLogicTablesRequestSyncStatus{
			value: "SUMMARY_FAILED",
		},
	}
}

func (c ListAggregationLogicTablesRequestSyncStatus) Value() string {
	return c.value
}

func (c ListAggregationLogicTablesRequestSyncStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAggregationLogicTablesRequestSyncStatus) UnmarshalJSON(b []byte) error {
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

type ListAggregationLogicTablesRequestSyncKey struct {
	value string
}

type ListAggregationLogicTablesRequestSyncKeyEnum struct {
	BUSINESS_ASSET  ListAggregationLogicTablesRequestSyncKey
	DATA_QUALITY    ListAggregationLogicTablesRequestSyncKey
	TECHNICAL_ASSET ListAggregationLogicTablesRequestSyncKey
	META_DATA_LINK  ListAggregationLogicTablesRequestSyncKey
	PHYSICAL_TABLE  ListAggregationLogicTablesRequestSyncKey
	DLF_TASK        ListAggregationLogicTablesRequestSyncKey
	MATERIALIZATION ListAggregationLogicTablesRequestSyncKey
}

func GetListAggregationLogicTablesRequestSyncKeyEnum() ListAggregationLogicTablesRequestSyncKeyEnum {
	return ListAggregationLogicTablesRequestSyncKeyEnum{
		BUSINESS_ASSET: ListAggregationLogicTablesRequestSyncKey{
			value: "BUSINESS_ASSET",
		},
		DATA_QUALITY: ListAggregationLogicTablesRequestSyncKey{
			value: "DATA_QUALITY",
		},
		TECHNICAL_ASSET: ListAggregationLogicTablesRequestSyncKey{
			value: "TECHNICAL_ASSET",
		},
		META_DATA_LINK: ListAggregationLogicTablesRequestSyncKey{
			value: "META_DATA_LINK",
		},
		PHYSICAL_TABLE: ListAggregationLogicTablesRequestSyncKey{
			value: "PHYSICAL_TABLE",
		},
		DLF_TASK: ListAggregationLogicTablesRequestSyncKey{
			value: "DLF_TASK",
		},
		MATERIALIZATION: ListAggregationLogicTablesRequestSyncKey{
			value: "MATERIALIZATION",
		},
	}
}

func (c ListAggregationLogicTablesRequestSyncKey) Value() string {
	return c.value
}

func (c ListAggregationLogicTablesRequestSyncKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAggregationLogicTablesRequestSyncKey) UnmarshalJSON(b []byte) error {
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
