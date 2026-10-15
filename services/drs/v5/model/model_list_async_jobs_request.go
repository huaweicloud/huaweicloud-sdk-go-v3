package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAsyncJobsRequest Request Object
type ListAsyncJobsRequest struct {

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us：英文 - zh-cn：中文  **默认取值**：  en-us
	XLanguage *ListAsyncJobsRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**：  批量异步创建的任务ID，用于按任务ID筛选查询结果。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	AsyncJobId *string `json:"async_job_id,omitempty"`

	// **参数解释**：  批量异步创建的任务状态，用于按任务状态筛选查询结果。  **约束限制**：  不涉及。  **取值范围**：  - ASYNC_JOB_VALIDATING：批量异步任务参数校验中 - ASYNC_JOB_VALIDATE_FAILED：批量异步任务参数校验失败 - AUTO_PARAM_VALIDATE_SUCCESS：批量异步任务参数校验成功 - COMMIT_SUCCESS：批量异步任务提交成功  **默认取值**：  不涉及
	Status *string `json:"status,omitempty"`

	// **参数解释**：  批量异步创建任务的租户名，用于按租户名筛选查询结果。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	DomainName *string `json:"domain_name,omitempty"`

	// **参数解释**：  批量异步创建任务的用户名，用于按用户名筛选查询结果。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**：  偏移量，表示查询该偏移量后面的记录，用于分页查询。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**：  查询返回记录的数量限制，用于分页查询时控制每页返回的数据条数。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：  返回结果按该关键字排序，指定批量异步任务列表的排序依据。  **约束限制**：  不涉及。  **取值范围**：  - create_time：按创建时间排序  **默认取值**：  create_time
	SortKey *string `json:"sort_key,omitempty"`

	// **参数解释**：  排序方向，指定返回结果按升序或降序排列。  **约束限制**：  不涉及。  **取值范围**：  - desc：降序 - asc：升序  **默认取值**：  desc
	SortDir *string `json:"sort_dir,omitempty"`
}

func (o ListAsyncJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAsyncJobsRequest struct{}"
	}

	return strings.Join([]string{"ListAsyncJobsRequest", string(data)}, " ")
}

type ListAsyncJobsRequestXLanguage struct {
	value string
}

type ListAsyncJobsRequestXLanguageEnum struct {
	EN_US ListAsyncJobsRequestXLanguage
	ZH_CN ListAsyncJobsRequestXLanguage
}

func GetListAsyncJobsRequestXLanguageEnum() ListAsyncJobsRequestXLanguageEnum {
	return ListAsyncJobsRequestXLanguageEnum{
		EN_US: ListAsyncJobsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ListAsyncJobsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ListAsyncJobsRequestXLanguage) Value() string {
	return c.value
}

func (c ListAsyncJobsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAsyncJobsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
