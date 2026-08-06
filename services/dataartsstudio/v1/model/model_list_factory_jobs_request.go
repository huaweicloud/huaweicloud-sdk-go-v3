package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListFactoryJobsRequest Request Object
type ListFactoryJobsRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 分页参数：每页限定数量
	Limit *int32 `json:"limit,omitempty"`

	// 分页参数：页数
	Offset *int32 `json:"offset,omitempty"`

	// 作业类型:  - REAL_TIME: 实时处理  - BATCH: 批处理
	JobType *ListFactoryJobsRequestJobType `json:"job_type,omitempty"`

	// 作业名称
	JobName *string `json:"job_name,omitempty"`

	// 作业ID，支持多个ID逗号分隔查询，最多50个ID，总长度不超过1000字符。 每个ID必须为纯数字。
	JobId *string `json:"job_id,omitempty"`

	// 作业状态，支持多个状态逗号分隔查询。 批处理作业状态：  - SCHEDULING: 调度中  - STOPPED: 停止  - PAUSED: 暂停 实时作业状态：  - STARTING: 启动中  - NORMAL: 正常  - EXCEPTION: 异常  - STOPPING: 停止中  - STOPPED: 停止  - PAUSE: 暂停  - ABNORMAL: 异常
	Status *string `json:"status,omitempty"`

	// 是否返回作业告警信息，默认为false。
	NeedAlarms *bool `json:"need_alarms,omitempty"`

	// 作业标签，多个标签逗号分隔。
	Tags *string `json:"tags,omitempty"`

	// 标签匹配模式：  - false: 任一标签匹配即返回（OR模式）  - true: 所有标签都匹配才返回（AND模式）
	MatchAllTags *bool `json:"match_all_tags,omitempty"`

	// 数据连接名称，按数据连接筛选作业。
	ConnectionName *string `json:"connection_name,omitempty"`

	// 源端数据连接类型，按源端数据类型筛选作业。
	SourceType *string `json:"source_type,omitempty"`

	// 源端数据连接名称，按源端数据名称筛选作业。
	SourceName *string `json:"source_name,omitempty"`

	// 目的端数据连接类型，按目的端数据类型筛选作业。
	SinkType *string `json:"sink_type,omitempty"`

	// 目的端数据连接名称，按目的端数据名称筛选作业。
	SinkName *string `json:"sink_name,omitempty"`
}

func (o ListFactoryJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFactoryJobsRequest struct{}"
	}

	return strings.Join([]string{"ListFactoryJobsRequest", string(data)}, " ")
}

type ListFactoryJobsRequestJobType struct {
	value string
}

type ListFactoryJobsRequestJobTypeEnum struct {
	REAL_TIME ListFactoryJobsRequestJobType
	BATCH     ListFactoryJobsRequestJobType
}

func GetListFactoryJobsRequestJobTypeEnum() ListFactoryJobsRequestJobTypeEnum {
	return ListFactoryJobsRequestJobTypeEnum{
		REAL_TIME: ListFactoryJobsRequestJobType{
			value: "REAL_TIME",
		},
		BATCH: ListFactoryJobsRequestJobType{
			value: "BATCH",
		},
	}
}

func (c ListFactoryJobsRequestJobType) Value() string {
	return c.value
}

func (c ListFactoryJobsRequestJobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFactoryJobsRequestJobType) UnmarshalJSON(b []byte) error {
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
