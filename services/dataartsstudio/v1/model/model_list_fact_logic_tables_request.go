package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListFactLogicTablesRequest Request Object
type ListFactLogicTablesRequest struct {

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
	Status *ListFactLogicTablesRequestStatus `json:"status,omitempty"`

	SyncStatus *ListFactLogicTablesRequestSyncStatus `json:"sync_status,omitempty"`

	SyncKey *[]ListFactLogicTablesRequestSyncKey `json:"sync_key,omitempty"`

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
}

func (o ListFactLogicTablesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFactLogicTablesRequest struct{}"
	}

	return strings.Join([]string{"ListFactLogicTablesRequest", string(data)}, " ")
}

type ListFactLogicTablesRequestStatus struct {
	value string
}

type ListFactLogicTablesRequestStatusEnum struct {
	DRAFT              ListFactLogicTablesRequestStatus
	PUBLISH_DEVELOPING ListFactLogicTablesRequestStatus
	PUBLISHED          ListFactLogicTablesRequestStatus
	OFFLINE_DEVELOPING ListFactLogicTablesRequestStatus
	OFFLINE            ListFactLogicTablesRequestStatus
	REJECT             ListFactLogicTablesRequestStatus
}

func GetListFactLogicTablesRequestStatusEnum() ListFactLogicTablesRequestStatusEnum {
	return ListFactLogicTablesRequestStatusEnum{
		DRAFT: ListFactLogicTablesRequestStatus{
			value: "DRAFT",
		},
		PUBLISH_DEVELOPING: ListFactLogicTablesRequestStatus{
			value: "PUBLISH_DEVELOPING",
		},
		PUBLISHED: ListFactLogicTablesRequestStatus{
			value: "PUBLISHED",
		},
		OFFLINE_DEVELOPING: ListFactLogicTablesRequestStatus{
			value: "OFFLINE_DEVELOPING",
		},
		OFFLINE: ListFactLogicTablesRequestStatus{
			value: "OFFLINE",
		},
		REJECT: ListFactLogicTablesRequestStatus{
			value: "REJECT",
		},
	}
}

func (c ListFactLogicTablesRequestStatus) Value() string {
	return c.value
}

func (c ListFactLogicTablesRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFactLogicTablesRequestStatus) UnmarshalJSON(b []byte) error {
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

type ListFactLogicTablesRequestSyncStatus struct {
	value string
}

type ListFactLogicTablesRequestSyncStatusEnum struct {
	RUNNING         ListFactLogicTablesRequestSyncStatus
	NO_NEED         ListFactLogicTablesRequestSyncStatus
	SUMMARY_SUCCESS ListFactLogicTablesRequestSyncStatus
	SUMMARY_FAILED  ListFactLogicTablesRequestSyncStatus
}

func GetListFactLogicTablesRequestSyncStatusEnum() ListFactLogicTablesRequestSyncStatusEnum {
	return ListFactLogicTablesRequestSyncStatusEnum{
		RUNNING: ListFactLogicTablesRequestSyncStatus{
			value: "RUNNING",
		},
		NO_NEED: ListFactLogicTablesRequestSyncStatus{
			value: "NO_NEED",
		},
		SUMMARY_SUCCESS: ListFactLogicTablesRequestSyncStatus{
			value: "SUMMARY_SUCCESS",
		},
		SUMMARY_FAILED: ListFactLogicTablesRequestSyncStatus{
			value: "SUMMARY_FAILED",
		},
	}
}

func (c ListFactLogicTablesRequestSyncStatus) Value() string {
	return c.value
}

func (c ListFactLogicTablesRequestSyncStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFactLogicTablesRequestSyncStatus) UnmarshalJSON(b []byte) error {
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

type ListFactLogicTablesRequestSyncKey struct {
	value string
}

type ListFactLogicTablesRequestSyncKeyEnum struct {
	BUSINESS_ASSET  ListFactLogicTablesRequestSyncKey
	DATA_QUALITY    ListFactLogicTablesRequestSyncKey
	TECHNICAL_ASSET ListFactLogicTablesRequestSyncKey
	META_DATA_LINK  ListFactLogicTablesRequestSyncKey
	PHYSICAL_TABLE  ListFactLogicTablesRequestSyncKey
	DLF_TASK        ListFactLogicTablesRequestSyncKey
	MATERIALIZATION ListFactLogicTablesRequestSyncKey
}

func GetListFactLogicTablesRequestSyncKeyEnum() ListFactLogicTablesRequestSyncKeyEnum {
	return ListFactLogicTablesRequestSyncKeyEnum{
		BUSINESS_ASSET: ListFactLogicTablesRequestSyncKey{
			value: "BUSINESS_ASSET",
		},
		DATA_QUALITY: ListFactLogicTablesRequestSyncKey{
			value: "DATA_QUALITY",
		},
		TECHNICAL_ASSET: ListFactLogicTablesRequestSyncKey{
			value: "TECHNICAL_ASSET",
		},
		META_DATA_LINK: ListFactLogicTablesRequestSyncKey{
			value: "META_DATA_LINK",
		},
		PHYSICAL_TABLE: ListFactLogicTablesRequestSyncKey{
			value: "PHYSICAL_TABLE",
		},
		DLF_TASK: ListFactLogicTablesRequestSyncKey{
			value: "DLF_TASK",
		},
		MATERIALIZATION: ListFactLogicTablesRequestSyncKey{
			value: "MATERIALIZATION",
		},
	}
}

func (c ListFactLogicTablesRequestSyncKey) Value() string {
	return c.value
}

func (c ListFactLogicTablesRequestSyncKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFactLogicTablesRequestSyncKey) UnmarshalJSON(b []byte) error {
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
