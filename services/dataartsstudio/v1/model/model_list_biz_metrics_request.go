package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListBizMetricsRequest Request Object
type ListBizMetricsRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 按名称或编码模糊查询
	Name *string `json:"name,omitempty"`

	// 按创建者查询
	CreateBy *string `json:"create_by,omitempty"`

	// 按负责人查询
	Owner *string `json:"owner,omitempty"`

	// 业务状态
	Status *ListBizMetricsRequestStatus `json:"status,omitempty"`

	SyncStatus *ListBizMetricsRequestSyncStatus `json:"sync_status,omitempty"`

	SyncKey *[]ListBizMetricsRequestSyncKey `json:"sync_key,omitempty"`

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
}

func (o ListBizMetricsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBizMetricsRequest struct{}"
	}

	return strings.Join([]string{"ListBizMetricsRequest", string(data)}, " ")
}

type ListBizMetricsRequestStatus struct {
	value string
}

type ListBizMetricsRequestStatusEnum struct {
	DRAFT              ListBizMetricsRequestStatus
	PUBLISH_DEVELOPING ListBizMetricsRequestStatus
	PUBLISHED          ListBizMetricsRequestStatus
	OFFLINE_DEVELOPING ListBizMetricsRequestStatus
	OFFLINE            ListBizMetricsRequestStatus
	REJECT             ListBizMetricsRequestStatus
}

func GetListBizMetricsRequestStatusEnum() ListBizMetricsRequestStatusEnum {
	return ListBizMetricsRequestStatusEnum{
		DRAFT: ListBizMetricsRequestStatus{
			value: "DRAFT",
		},
		PUBLISH_DEVELOPING: ListBizMetricsRequestStatus{
			value: "PUBLISH_DEVELOPING",
		},
		PUBLISHED: ListBizMetricsRequestStatus{
			value: "PUBLISHED",
		},
		OFFLINE_DEVELOPING: ListBizMetricsRequestStatus{
			value: "OFFLINE_DEVELOPING",
		},
		OFFLINE: ListBizMetricsRequestStatus{
			value: "OFFLINE",
		},
		REJECT: ListBizMetricsRequestStatus{
			value: "REJECT",
		},
	}
}

func (c ListBizMetricsRequestStatus) Value() string {
	return c.value
}

func (c ListBizMetricsRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListBizMetricsRequestStatus) UnmarshalJSON(b []byte) error {
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

type ListBizMetricsRequestSyncStatus struct {
	value string
}

type ListBizMetricsRequestSyncStatusEnum struct {
	RUNNING         ListBizMetricsRequestSyncStatus
	NO_NEED         ListBizMetricsRequestSyncStatus
	SUMMARY_SUCCESS ListBizMetricsRequestSyncStatus
	SUMMARY_FAILED  ListBizMetricsRequestSyncStatus
}

func GetListBizMetricsRequestSyncStatusEnum() ListBizMetricsRequestSyncStatusEnum {
	return ListBizMetricsRequestSyncStatusEnum{
		RUNNING: ListBizMetricsRequestSyncStatus{
			value: "RUNNING",
		},
		NO_NEED: ListBizMetricsRequestSyncStatus{
			value: "NO_NEED",
		},
		SUMMARY_SUCCESS: ListBizMetricsRequestSyncStatus{
			value: "SUMMARY_SUCCESS",
		},
		SUMMARY_FAILED: ListBizMetricsRequestSyncStatus{
			value: "SUMMARY_FAILED",
		},
	}
}

func (c ListBizMetricsRequestSyncStatus) Value() string {
	return c.value
}

func (c ListBizMetricsRequestSyncStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListBizMetricsRequestSyncStatus) UnmarshalJSON(b []byte) error {
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

type ListBizMetricsRequestSyncKey struct {
	value string
}

type ListBizMetricsRequestSyncKeyEnum struct {
	BUSINESS_ASSET  ListBizMetricsRequestSyncKey
	DATA_QUALITY    ListBizMetricsRequestSyncKey
	TECHNICAL_ASSET ListBizMetricsRequestSyncKey
	META_DATA_LINK  ListBizMetricsRequestSyncKey
	PHYSICAL_TABLE  ListBizMetricsRequestSyncKey
	DLF_TASK        ListBizMetricsRequestSyncKey
	MATERIALIZATION ListBizMetricsRequestSyncKey
}

func GetListBizMetricsRequestSyncKeyEnum() ListBizMetricsRequestSyncKeyEnum {
	return ListBizMetricsRequestSyncKeyEnum{
		BUSINESS_ASSET: ListBizMetricsRequestSyncKey{
			value: "BUSINESS_ASSET",
		},
		DATA_QUALITY: ListBizMetricsRequestSyncKey{
			value: "DATA_QUALITY",
		},
		TECHNICAL_ASSET: ListBizMetricsRequestSyncKey{
			value: "TECHNICAL_ASSET",
		},
		META_DATA_LINK: ListBizMetricsRequestSyncKey{
			value: "META_DATA_LINK",
		},
		PHYSICAL_TABLE: ListBizMetricsRequestSyncKey{
			value: "PHYSICAL_TABLE",
		},
		DLF_TASK: ListBizMetricsRequestSyncKey{
			value: "DLF_TASK",
		},
		MATERIALIZATION: ListBizMetricsRequestSyncKey{
			value: "MATERIALIZATION",
		},
	}
}

func (c ListBizMetricsRequestSyncKey) Value() string {
	return c.value
}

func (c ListBizMetricsRequestSyncKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListBizMetricsRequestSyncKey) UnmarshalJSON(b []byte) error {
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
