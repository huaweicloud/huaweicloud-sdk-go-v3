package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAllTablesRequest Request Object
type ListAllTablesRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 按名称或编码模糊查询
	Name *string `json:"name,omitempty"`

	// 按创建者查询
	CreateBy *string `json:"create_by,omitempty"`

	// 业务状态
	Status *ListAllTablesRequestStatus `json:"status,omitempty"`

	SyncStatus *ListAllTablesRequestSyncStatus `json:"sync_status,omitempty"`

	SyncKey *[]ListAllTablesRequestSyncKey `json:"sync_key,omitempty"`

	// 所属的业务分层的id
	BizCatalogId *string `json:"biz_catalog_id,omitempty"`

	// 时间过滤左边界,与end_time一起使用,只支持时间范围过滤,单边过滤无效
	BeginTime *string `json:"begin_time,omitempty"`

	// 时间过滤右边界,与begin_time一起使用只支持时间范围过滤,单边过滤无效
	EndTime *string `json:"end_time,omitempty"`

	// 查询条数，即查询Y条数据。默认值50，取值范围[1,100]
	Limit *int32 `json:"limit,omitempty"`

	// 查询起始坐标，即跳过X条数据，仅支持0或limit的整数倍，不满足则向下取整。默认值0
	Offset *int32 `json:"offset,omitempty"`

	// 所属主题的id列表
	BizCatalogIdList *[]int64 `json:"biz_catalog_id_list,omitempty"`

	// 查询的表类型，必填
	BizTypeList []ListAllTablesRequestBizTypeList `json:"biz_type_list"`
}

func (o ListAllTablesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllTablesRequest struct{}"
	}

	return strings.Join([]string{"ListAllTablesRequest", string(data)}, " ")
}

type ListAllTablesRequestStatus struct {
	value string
}

type ListAllTablesRequestStatusEnum struct {
	DRAFT              ListAllTablesRequestStatus
	PUBLISH_DEVELOPING ListAllTablesRequestStatus
	PUBLISHED          ListAllTablesRequestStatus
	OFFLINE_DEVELOPING ListAllTablesRequestStatus
	OFFLINE            ListAllTablesRequestStatus
	REJECT             ListAllTablesRequestStatus
}

func GetListAllTablesRequestStatusEnum() ListAllTablesRequestStatusEnum {
	return ListAllTablesRequestStatusEnum{
		DRAFT: ListAllTablesRequestStatus{
			value: "DRAFT",
		},
		PUBLISH_DEVELOPING: ListAllTablesRequestStatus{
			value: "PUBLISH_DEVELOPING",
		},
		PUBLISHED: ListAllTablesRequestStatus{
			value: "PUBLISHED",
		},
		OFFLINE_DEVELOPING: ListAllTablesRequestStatus{
			value: "OFFLINE_DEVELOPING",
		},
		OFFLINE: ListAllTablesRequestStatus{
			value: "OFFLINE",
		},
		REJECT: ListAllTablesRequestStatus{
			value: "REJECT",
		},
	}
}

func (c ListAllTablesRequestStatus) Value() string {
	return c.value
}

func (c ListAllTablesRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAllTablesRequestStatus) UnmarshalJSON(b []byte) error {
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

type ListAllTablesRequestSyncStatus struct {
	value string
}

type ListAllTablesRequestSyncStatusEnum struct {
	RUNNING         ListAllTablesRequestSyncStatus
	NO_NEED         ListAllTablesRequestSyncStatus
	SUMMARY_SUCCESS ListAllTablesRequestSyncStatus
	SUMMARY_FAILED  ListAllTablesRequestSyncStatus
}

func GetListAllTablesRequestSyncStatusEnum() ListAllTablesRequestSyncStatusEnum {
	return ListAllTablesRequestSyncStatusEnum{
		RUNNING: ListAllTablesRequestSyncStatus{
			value: "RUNNING",
		},
		NO_NEED: ListAllTablesRequestSyncStatus{
			value: "NO_NEED",
		},
		SUMMARY_SUCCESS: ListAllTablesRequestSyncStatus{
			value: "SUMMARY_SUCCESS",
		},
		SUMMARY_FAILED: ListAllTablesRequestSyncStatus{
			value: "SUMMARY_FAILED",
		},
	}
}

func (c ListAllTablesRequestSyncStatus) Value() string {
	return c.value
}

func (c ListAllTablesRequestSyncStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAllTablesRequestSyncStatus) UnmarshalJSON(b []byte) error {
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

type ListAllTablesRequestSyncKey struct {
	value string
}

type ListAllTablesRequestSyncKeyEnum struct {
	BUSINESS_ASSET  ListAllTablesRequestSyncKey
	DATA_QUALITY    ListAllTablesRequestSyncKey
	TECHNICAL_ASSET ListAllTablesRequestSyncKey
	META_DATA_LINK  ListAllTablesRequestSyncKey
	PHYSICAL_TABLE  ListAllTablesRequestSyncKey
	DLF_TASK        ListAllTablesRequestSyncKey
	MATERIALIZATION ListAllTablesRequestSyncKey
}

func GetListAllTablesRequestSyncKeyEnum() ListAllTablesRequestSyncKeyEnum {
	return ListAllTablesRequestSyncKeyEnum{
		BUSINESS_ASSET: ListAllTablesRequestSyncKey{
			value: "BUSINESS_ASSET",
		},
		DATA_QUALITY: ListAllTablesRequestSyncKey{
			value: "DATA_QUALITY",
		},
		TECHNICAL_ASSET: ListAllTablesRequestSyncKey{
			value: "TECHNICAL_ASSET",
		},
		META_DATA_LINK: ListAllTablesRequestSyncKey{
			value: "META_DATA_LINK",
		},
		PHYSICAL_TABLE: ListAllTablesRequestSyncKey{
			value: "PHYSICAL_TABLE",
		},
		DLF_TASK: ListAllTablesRequestSyncKey{
			value: "DLF_TASK",
		},
		MATERIALIZATION: ListAllTablesRequestSyncKey{
			value: "MATERIALIZATION",
		},
	}
}

func (c ListAllTablesRequestSyncKey) Value() string {
	return c.value
}

func (c ListAllTablesRequestSyncKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAllTablesRequestSyncKey) UnmarshalJSON(b []byte) error {
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

type ListAllTablesRequestBizTypeList struct {
	value string
}

type ListAllTablesRequestBizTypeListEnum struct {
	TABLE_MODEL_LOGIC       ListAllTablesRequestBizTypeList
	TABLE_MODEL             ListAllTablesRequestBizTypeList
	DIMENSION_LOGIC_TABLE   ListAllTablesRequestBizTypeList
	FACT_LOGIC_TABLE        ListAllTablesRequestBizTypeList
	AGGREGATION_LOGIC_TABLE ListAllTablesRequestBizTypeList
}

func GetListAllTablesRequestBizTypeListEnum() ListAllTablesRequestBizTypeListEnum {
	return ListAllTablesRequestBizTypeListEnum{
		TABLE_MODEL_LOGIC: ListAllTablesRequestBizTypeList{
			value: "TABLE_MODEL_LOGIC",
		},
		TABLE_MODEL: ListAllTablesRequestBizTypeList{
			value: "TABLE_MODEL",
		},
		DIMENSION_LOGIC_TABLE: ListAllTablesRequestBizTypeList{
			value: "DIMENSION_LOGIC_TABLE",
		},
		FACT_LOGIC_TABLE: ListAllTablesRequestBizTypeList{
			value: "FACT_LOGIC_TABLE",
		},
		AGGREGATION_LOGIC_TABLE: ListAllTablesRequestBizTypeList{
			value: "AGGREGATION_LOGIC_TABLE",
		},
	}
}

func (c ListAllTablesRequestBizTypeList) Value() string {
	return c.value
}

func (c ListAllTablesRequestBizTypeList) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAllTablesRequestBizTypeList) UnmarshalJSON(b []byte) error {
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
