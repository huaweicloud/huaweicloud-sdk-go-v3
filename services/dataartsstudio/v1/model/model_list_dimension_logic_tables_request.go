package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListDimensionLogicTablesRequest Request Object
type ListDimensionLogicTablesRequest struct {

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

	// 按审核人查询。
	Approver *string `json:"approver,omitempty"`

	// 按负责人查询。
	Owner *string `json:"owner,omitempty"`

	// 业务状态。 枚举值：   - DRAFT: 草稿   - PUBLISH_DEVELOPING: 发布待审批   - PUBLISHED: 已发布   - OFFLINE_DEVELOPING: 下线待审批   - OFFLINE: 已下线   - REJECT: 已驳回
	Status *ListDimensionLogicTablesRequestStatus `json:"status,omitempty"`

	// 同步状态枚举。 枚举值：   - RUNNING: 同步中   - NO_NEED: 未同步   - SUMMARY_SUCCESS: 整体成功   - SUMMARY_FAILED: 整体失败
	SyncStatus *ListDimensionLogicTablesRequestSyncStatus `json:"sync_status,omitempty"`

	// 同步任务类型枚举。 枚举值：   - BUSINESS_ASSET: 同步业务资产   - DATA_QUALITY: 创建质量作业   - TECHNICAL_ASSET: 同步技术资产   - META_DATA_LINK: 资产关联   - PHYSICAL_TABLE: 创建表（生产环境）   - DEV_PHYSICAL_TABLE: 创建表（开发环境）   - DLF_TASK: 创建数据开发作业   - MATERIALIZATION: 数值落库（码表）   - PUBLISH_TO_DLM: 发布数据服务API   - SUMMARY_STATUS: 整体状态
	SyncKey *[]ListDimensionLogicTablesRequestSyncKey `json:"sync_key,omitempty"`

	// 主题域l2的ID，填写String类型替代Long类型。
	L2Id *string `json:"l2_id,omitempty"`

	// 依据维度ID查维度属性，填写String类型替代Long类型。
	DimensionId *string `json:"dimension_id,omitempty"`

	// 时间过滤左边界，与end_time一起使用，只支持时间范围过滤，单边过滤无效。格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	BeginTime *string `json:"begin_time,omitempty"`

	// 时间过滤右边界，与begin_time一起使用只支持时间范围过滤，单边过滤无效。格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	EndTime *string `json:"end_time,omitempty"`

	// 每页查询条数，即查询Y条数据。默认值50，取值范围[1,100]。
	Limit *int32 `json:"limit,omitempty"`

	// 查询起始坐标，即跳过X条数据，仅支持0或limit的整数倍，不满足则向下取整，默认值0。
	Offset *int32 `json:"offset,omitempty"`

	// 所属的业务分层的ID。
	BizCatalogId *string `json:"biz_catalog_id,omitempty"`

	// 维度类型。 枚举值：   - COMMON: 普通维度   - LOOKUP: 码表维度   - HIERARCHIES: 层级维度
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
	BUSINESS_ASSET     ListDimensionLogicTablesRequestSyncKey
	DATA_QUALITY       ListDimensionLogicTablesRequestSyncKey
	TECHNICAL_ASSET    ListDimensionLogicTablesRequestSyncKey
	META_DATA_LINK     ListDimensionLogicTablesRequestSyncKey
	PHYSICAL_TABLE     ListDimensionLogicTablesRequestSyncKey
	DEV_PHYSICAL_TABLE ListDimensionLogicTablesRequestSyncKey
	DLF_TASK           ListDimensionLogicTablesRequestSyncKey
	MATERIALIZATION    ListDimensionLogicTablesRequestSyncKey
	PUBLISH_TO_DLM     ListDimensionLogicTablesRequestSyncKey
	SUMMARY_STATUS     ListDimensionLogicTablesRequestSyncKey
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
		DEV_PHYSICAL_TABLE: ListDimensionLogicTablesRequestSyncKey{
			value: "DEV_PHYSICAL_TABLE",
		},
		DLF_TASK: ListDimensionLogicTablesRequestSyncKey{
			value: "DLF_TASK",
		},
		MATERIALIZATION: ListDimensionLogicTablesRequestSyncKey{
			value: "MATERIALIZATION",
		},
		PUBLISH_TO_DLM: ListDimensionLogicTablesRequestSyncKey{
			value: "PUBLISH_TO_DLM",
		},
		SUMMARY_STATUS: ListDimensionLogicTablesRequestSyncKey{
			value: "SUMMARY_STATUS",
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
