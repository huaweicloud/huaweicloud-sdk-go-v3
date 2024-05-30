package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListBizMetricsRequest Request Object
type ListBizMetricsRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 项目ID，获取方法请参见[项目ID和账号ID](projectid_accountid.xml)。  多project场景采用AK/SK认证的接口请求，则该字段必选。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 默认值：application/json;charset=UTF-8 可选，有Body体的情况下必选，没有Body体则无需填写和校验。
	ContentType *string `json:"Content-Type,omitempty"`

	// 按名称或编码模糊查询。
	Name *string `json:"name,omitempty"`

	// 按创建者查询。
	CreateBy *string `json:"create_by,omitempty"`

	// 按负责人查询。
	Owner *string `json:"owner,omitempty"`

	// 业务状态。 枚举值：   - DRAFT: 草稿   - PUBLISH_DEVELOPING: 发布待审批   - PUBLISHED: 已发布   - OFFLINE_DEVELOPING: 下线待审批   - OFFLINE: 已下线   - REJECT: 已驳回
	Status *ListBizMetricsRequestStatus `json:"status,omitempty"`

	// 同步状态枚举。 枚举值：   - RUNNING: 同步中   - NO_NEED: 未同步   - SUMMARY_SUCCESS: 整体成功   - SUMMARY_FAILED: 整体失败
	SyncStatus *ListBizMetricsRequestSyncStatus `json:"sync_status,omitempty"`

	// 同步任务类型枚举。 枚举值：   - BUSINESS_ASSET: 同步业务资产   - DATA_QUALITY: 创建质量作业   - TECHNICAL_ASSET: 同步技术资产   - META_DATA_LINK: 资产关联   - PHYSICAL_TABLE: 创建表（生产环境）   - DEV_PHYSICAL_TABLE: 创建表（开发环境）   - DLF_TASK: 创建数据开发作业   - MATERIALIZATION: 数值落库（码表）   - PUBLISH_TO_DLM: 发布数据服务API   - SUMMARY_STATUS: 整体状态
	SyncKey *[]ListBizMetricsRequestSyncKey `json:"sync_key,omitempty"`

	// 所属的业务分层的ID。
	BizCatalogId *string `json:"biz_catalog_id,omitempty"`

	// 时间过滤左边界，与end_time一起使用，只支持时间范围过滤，单边过滤无效。格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	BeginTime *string `json:"begin_time,omitempty"`

	// 时间过滤右边界，与begin_time一起使用只支持时间范围过滤，单边过滤无效。格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	EndTime *string `json:"end_time,omitempty"`

	// 每页查询条数，即查询Y条数据。默认值50，取值范围[1,100]。
	Limit *int32 `json:"limit,omitempty"`

	// 查询起始坐标，即跳过X条数据，仅支持0或limit的整数倍，不满足则向下取整，默认值0。
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
	BUSINESS_ASSET     ListBizMetricsRequestSyncKey
	DATA_QUALITY       ListBizMetricsRequestSyncKey
	TECHNICAL_ASSET    ListBizMetricsRequestSyncKey
	META_DATA_LINK     ListBizMetricsRequestSyncKey
	PHYSICAL_TABLE     ListBizMetricsRequestSyncKey
	DEV_PHYSICAL_TABLE ListBizMetricsRequestSyncKey
	DLF_TASK           ListBizMetricsRequestSyncKey
	MATERIALIZATION    ListBizMetricsRequestSyncKey
	PUBLISH_TO_DLM     ListBizMetricsRequestSyncKey
	SUMMARY_STATUS     ListBizMetricsRequestSyncKey
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
		DEV_PHYSICAL_TABLE: ListBizMetricsRequestSyncKey{
			value: "DEV_PHYSICAL_TABLE",
		},
		DLF_TASK: ListBizMetricsRequestSyncKey{
			value: "DLF_TASK",
		},
		MATERIALIZATION: ListBizMetricsRequestSyncKey{
			value: "MATERIALIZATION",
		},
		PUBLISH_TO_DLM: ListBizMetricsRequestSyncKey{
			value: "PUBLISH_TO_DLM",
		},
		SUMMARY_STATUS: ListBizMetricsRequestSyncKey{
			value: "SUMMARY_STATUS",
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
