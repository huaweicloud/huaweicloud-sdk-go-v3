package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListReplicationJobsRequest Request Object
type ListReplicationJobsRequest struct {

	// 请求语言类型。 en-us：英文 zh-cn：中文
	XLanguage *ListReplicationJobsRequestXLanguage `json:"X-Language,omitempty"`

	// 任务名称，支持模糊搜索。
	Name *string `json:"name,omitempty"`

	// 备份迁移任务状态。 TRANSFERRING：恢复中 SUCCESS：成功 FAILED：失败 PRECHECK FAILED：预检查失败
	Status *ListReplicationJobsRequestStatus `json:"status,omitempty"`

	// 数据库实例ID，最大数量为10。
	DbsInstanceIds *[]string `json:"dbs_instance_ids,omitempty"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 创建时间。
	CreateAt *string `json:"create_at,omitempty"`

	// 完成时间。
	CompletedAt *string `json:"completed_at,omitempty"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 标签。
	Tags *string `json:"tags,omitempty"`

	// 查询返回记录的数量限制，默认值为10。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，默认值为0，表示查询该偏移量后面的记录。
	Offset *int32 `json:"offset,omitempty"`

	// 排序字段。
	SortKey *string `json:"sort_key,omitempty"`

	// 排序方法。
	SortDir *string `json:"sort_dir,omitempty"`
}

func (o ListReplicationJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListReplicationJobsRequest struct{}"
	}

	return strings.Join([]string{"ListReplicationJobsRequest", string(data)}, " ")
}

type ListReplicationJobsRequestXLanguage struct {
	value string
}

type ListReplicationJobsRequestXLanguageEnum struct {
	EN_US ListReplicationJobsRequestXLanguage
	ZH_CN ListReplicationJobsRequestXLanguage
}

func GetListReplicationJobsRequestXLanguageEnum() ListReplicationJobsRequestXLanguageEnum {
	return ListReplicationJobsRequestXLanguageEnum{
		EN_US: ListReplicationJobsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ListReplicationJobsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ListReplicationJobsRequestXLanguage) Value() string {
	return c.value
}

func (c ListReplicationJobsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListReplicationJobsRequestXLanguage) UnmarshalJSON(b []byte) error {
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

type ListReplicationJobsRequestStatus struct {
	value string
}

type ListReplicationJobsRequestStatusEnum struct {
	TRANSFERRING    ListReplicationJobsRequestStatus
	SUCCESS         ListReplicationJobsRequestStatus
	FAILED          ListReplicationJobsRequestStatus
	PRECHECK_FAILED ListReplicationJobsRequestStatus
}

func GetListReplicationJobsRequestStatusEnum() ListReplicationJobsRequestStatusEnum {
	return ListReplicationJobsRequestStatusEnum{
		TRANSFERRING: ListReplicationJobsRequestStatus{
			value: "TRANSFERRING",
		},
		SUCCESS: ListReplicationJobsRequestStatus{
			value: "SUCCESS",
		},
		FAILED: ListReplicationJobsRequestStatus{
			value: "FAILED",
		},
		PRECHECK_FAILED: ListReplicationJobsRequestStatus{
			value: "PRECHECK FAILED",
		},
	}
}

func (c ListReplicationJobsRequestStatus) Value() string {
	return c.value
}

func (c ListReplicationJobsRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListReplicationJobsRequestStatus) UnmarshalJSON(b []byte) error {
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
