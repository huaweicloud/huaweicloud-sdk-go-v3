package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAllTablesRequest Request Object
type ListAllTablesRequest struct {

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

	// 业务状态。 枚举值：   - DRAFT: 草稿   - PUBLISH_DEVELOPING: 发布待审批   - PUBLISHED: 已发布   - OFFLINE_DEVELOPING: 下线待审批   - OFFLINE: 已下线   - REJECT: 已驳回
	Status *ListAllTablesRequestStatus `json:"status,omitempty"`

	// 同步状态枚举。 枚举值：   - RUNNING: 同步中   - NO_NEED: 未同步   - SUMMARY_SUCCESS: 整体成功   - SUMMARY_FAILED: 整体失败
	SyncStatus *ListAllTablesRequestSyncStatus `json:"sync_status,omitempty"`

	// 同步任务类型枚举。 枚举值：   - BUSINESS_ASSET: 同步业务资产   - DATA_QUALITY: 创建质量作业   - TECHNICAL_ASSET: 同步技术资产   - META_DATA_LINK: 资产关联   - PHYSICAL_TABLE: 创建表（生产环境）   - DEV_PHYSICAL_TABLE: 创建表（开发环境）   - DLF_TASK: 创建数据开发作业   - MATERIALIZATION: 数值落库（码表）   - PUBLISH_TO_DLM: 发布数据服务API   - SUMMARY_STATUS: 整体状态
	SyncKey *[]ListAllTablesRequestSyncKey `json:"sync_key,omitempty"`

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

	// 所属主题的ID列表，填写String类型替代Long类型。
	BizCatalogIdList *[]string `json:"biz_catalog_id_list,omitempty"`

	// 查询的表类型，必填。 枚举值：   - TABLE_MODEL_LOGIC: 逻辑实体   - TABLE_MODEL: 物理表   - DIMENSION_LOGIC_TABLE: 维度表   - FACT_LOGIC_TABLE: 事实表   - AGGREGATION_LOGIC_TABLE: 汇总表
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
	BUSINESS_ASSET     ListAllTablesRequestSyncKey
	DATA_QUALITY       ListAllTablesRequestSyncKey
	TECHNICAL_ASSET    ListAllTablesRequestSyncKey
	META_DATA_LINK     ListAllTablesRequestSyncKey
	PHYSICAL_TABLE     ListAllTablesRequestSyncKey
	DEV_PHYSICAL_TABLE ListAllTablesRequestSyncKey
	DLF_TASK           ListAllTablesRequestSyncKey
	MATERIALIZATION    ListAllTablesRequestSyncKey
	PUBLISH_TO_DLM     ListAllTablesRequestSyncKey
	SUMMARY_STATUS     ListAllTablesRequestSyncKey
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
		DEV_PHYSICAL_TABLE: ListAllTablesRequestSyncKey{
			value: "DEV_PHYSICAL_TABLE",
		},
		DLF_TASK: ListAllTablesRequestSyncKey{
			value: "DLF_TASK",
		},
		MATERIALIZATION: ListAllTablesRequestSyncKey{
			value: "MATERIALIZATION",
		},
		PUBLISH_TO_DLM: ListAllTablesRequestSyncKey{
			value: "PUBLISH_TO_DLM",
		},
		SUMMARY_STATUS: ListAllTablesRequestSyncKey{
			value: "SUMMARY_STATUS",
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
