package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListTableModelsRequest Request Object
type ListTableModelsRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 按名称或编码模糊查询
	Name *string `json:"name,omitempty"`

	// 按创建者查询
	CreateBy *string `json:"create_by,omitempty"`

	// 按审核人查询
	Approver *string `json:"approver,omitempty"`

	// 业务状态
	Status *ListTableModelsRequestStatus `json:"status,omitempty"`

	SyncStatus *ListTableModelsRequestSyncStatus `json:"sync_status,omitempty"`

	SyncKey *[]ListTableModelsRequestSyncKey `json:"sync_key,omitempty"`

	// 时间过滤左边界,与end_time一起使用,只支持时间范围过滤,单边过滤无效
	BeginTime *string `json:"begin_time,omitempty"`

	// 时间过滤右边界,与begin_time一起使用只支持时间范围过滤,单边过滤无效
	EndTime *string `json:"end_time,omitempty"`

	// 查询条数，即查询Y条数据。默认值50，取值范围[1,100]
	Limit *int32 `json:"limit,omitempty"`

	// 查询起始坐标，即跳过X条数据，仅支持0或limit的整数倍，不满足则向下取整。默认值0
	Offset *int32 `json:"offset,omitempty"`

	// 所属的模型的id
	ModelId string `json:"model_id"`

	// 所属的业务分层的id
	BizCatalogId *string `json:"biz_catalog_id,omitempty"`
}

func (o ListTableModelsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTableModelsRequest struct{}"
	}

	return strings.Join([]string{"ListTableModelsRequest", string(data)}, " ")
}

type ListTableModelsRequestStatus struct {
	value string
}

type ListTableModelsRequestStatusEnum struct {
	DRAFT              ListTableModelsRequestStatus
	PUBLISH_DEVELOPING ListTableModelsRequestStatus
	PUBLISHED          ListTableModelsRequestStatus
	OFFLINE_DEVELOPING ListTableModelsRequestStatus
	OFFLINE            ListTableModelsRequestStatus
	REJECT             ListTableModelsRequestStatus
}

func GetListTableModelsRequestStatusEnum() ListTableModelsRequestStatusEnum {
	return ListTableModelsRequestStatusEnum{
		DRAFT: ListTableModelsRequestStatus{
			value: "DRAFT",
		},
		PUBLISH_DEVELOPING: ListTableModelsRequestStatus{
			value: "PUBLISH_DEVELOPING",
		},
		PUBLISHED: ListTableModelsRequestStatus{
			value: "PUBLISHED",
		},
		OFFLINE_DEVELOPING: ListTableModelsRequestStatus{
			value: "OFFLINE_DEVELOPING",
		},
		OFFLINE: ListTableModelsRequestStatus{
			value: "OFFLINE",
		},
		REJECT: ListTableModelsRequestStatus{
			value: "REJECT",
		},
	}
}

func (c ListTableModelsRequestStatus) Value() string {
	return c.value
}

func (c ListTableModelsRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTableModelsRequestStatus) UnmarshalJSON(b []byte) error {
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

type ListTableModelsRequestSyncStatus struct {
	value string
}

type ListTableModelsRequestSyncStatusEnum struct {
	RUNNING         ListTableModelsRequestSyncStatus
	NO_NEED         ListTableModelsRequestSyncStatus
	SUMMARY_SUCCESS ListTableModelsRequestSyncStatus
	SUMMARY_FAILED  ListTableModelsRequestSyncStatus
}

func GetListTableModelsRequestSyncStatusEnum() ListTableModelsRequestSyncStatusEnum {
	return ListTableModelsRequestSyncStatusEnum{
		RUNNING: ListTableModelsRequestSyncStatus{
			value: "RUNNING",
		},
		NO_NEED: ListTableModelsRequestSyncStatus{
			value: "NO_NEED",
		},
		SUMMARY_SUCCESS: ListTableModelsRequestSyncStatus{
			value: "SUMMARY_SUCCESS",
		},
		SUMMARY_FAILED: ListTableModelsRequestSyncStatus{
			value: "SUMMARY_FAILED",
		},
	}
}

func (c ListTableModelsRequestSyncStatus) Value() string {
	return c.value
}

func (c ListTableModelsRequestSyncStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTableModelsRequestSyncStatus) UnmarshalJSON(b []byte) error {
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

type ListTableModelsRequestSyncKey struct {
	value string
}

type ListTableModelsRequestSyncKeyEnum struct {
	BUSINESS_ASSET  ListTableModelsRequestSyncKey
	DATA_QUALITY    ListTableModelsRequestSyncKey
	TECHNICAL_ASSET ListTableModelsRequestSyncKey
	META_DATA_LINK  ListTableModelsRequestSyncKey
	PHYSICAL_TABLE  ListTableModelsRequestSyncKey
	DLF_TASK        ListTableModelsRequestSyncKey
	MATERIALIZATION ListTableModelsRequestSyncKey
}

func GetListTableModelsRequestSyncKeyEnum() ListTableModelsRequestSyncKeyEnum {
	return ListTableModelsRequestSyncKeyEnum{
		BUSINESS_ASSET: ListTableModelsRequestSyncKey{
			value: "BUSINESS_ASSET",
		},
		DATA_QUALITY: ListTableModelsRequestSyncKey{
			value: "DATA_QUALITY",
		},
		TECHNICAL_ASSET: ListTableModelsRequestSyncKey{
			value: "TECHNICAL_ASSET",
		},
		META_DATA_LINK: ListTableModelsRequestSyncKey{
			value: "META_DATA_LINK",
		},
		PHYSICAL_TABLE: ListTableModelsRequestSyncKey{
			value: "PHYSICAL_TABLE",
		},
		DLF_TASK: ListTableModelsRequestSyncKey{
			value: "DLF_TASK",
		},
		MATERIALIZATION: ListTableModelsRequestSyncKey{
			value: "MATERIALIZATION",
		},
	}
}

func (c ListTableModelsRequestSyncKey) Value() string {
	return c.value
}

func (c ListTableModelsRequestSyncKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTableModelsRequestSyncKey) UnmarshalJSON(b []byte) error {
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
