package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListFlinkJobsRequest Request Object
type ListFlinkJobsRequest struct {

	// 参数解释:  作业类型 示例: flink_jar_job 约束限制:  无 取值范围: flink_sql_job（flink sql作业） flink_opensource_sql_job（flink opensource sql作业） flink_sql_edge_job（flink sql边缘作业） flink_jar_job（flink自定义作业） 默认取值: 无
	JobType *ListFlinkJobsRequestJobType `json:"job_type,omitempty"`

	// 参数解释:  返回的数据条数。默认为10 示例: 100 约束限制:  无 取值范围: 大于等于0的整数 默认取值: 10
	Limit *int32 `json:"limit,omitempty"`

	// 参数解释:  作业名称 示例: myjob 约束限制:  长度在[0,57]的字符串 取值范围: 无 默认取值: 无
	Name *string `json:"name,omitempty"`

	// 参数解释:  作业偏移量 示例: 10000 约束限制:  无 取值范围: 大于等于0的整数 默认取值: 0
	Offset *int64 `json:"offset,omitempty"`

	// 参数解释:  查询结果排序 示例: asc 约束限制:  无 取值范围: asc desc 默认取值: desc
	Order *string `json:"order,omitempty"`

	// 参数解释:  队列名称 示例: queue1 约束限制:  无 取值范围: 无 默认取值: 无
	QueueName *string `json:"queue_name,omitempty"`

	// 参数解释:  边缘父作业ID, 用于查询指定边缘作业的子作业。不带该参数时, 查询所有非边缘作业和边缘父作业, 不包括边缘子作业 示例: 548483 约束限制:  无 取值范围: 无 默认取值: 无
	RootJobId *int64 `json:"root_job_id,omitempty"`

	// 参数解释:  是否返回作业详情信息 示例: false 约束限制:  无 取值范围: true,false 默认取值: false
	ShowDetail *bool `json:"show_detail,omitempty"`

	// 参数解释:  作业状态 示例: job_submitting 约束限制:  无 取值范围: job_init（草稿） job_submitting（提交中） job_submit_fail（提交失败） job_running（运行中） job_running_exception（运行异常） job_downloading（下载中） job_idle（空闲） job_canceling（停止中） job_cancel_success（已停止） job_cancel_fail（停止失败） job_savepointing（保存点创建中） job_arrearage_stopped（因欠费被停止） job_arrearage_recovering（欠费作业恢复中） job_finish（已完成） 默认取值: 无
	Status *ListFlinkJobsRequestStatus `json:"status,omitempty"`

	// 参数解释:  企业项目名称 示例: DLI 约束限制:  无 取值范围: 无 默认取值: 无
	SysEnterpriseProjectName *string `json:"sys_enterprise_project_name,omitempty"`

	// 参数解释:  标签列表 示例: key_zy1=zy01,AA=aa 约束限制:  符合键值对格式(如“key=value”)的字符串 取值范围: 无 默认取值: 无
	Tags *string `json:"tags,omitempty"`

	// 参数解释:  用户名，可作为筛选条件 示例: ei_dlics_d00352431 约束限制:  无 取值范围: 无 默认取值: 无
	UserName *string `json:"user_name,omitempty"`
}

func (o ListFlinkJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlinkJobsRequest struct{}"
	}

	return strings.Join([]string{"ListFlinkJobsRequest", string(data)}, " ")
}

type ListFlinkJobsRequestJobType struct {
	value string
}

type ListFlinkJobsRequestJobTypeEnum struct {
	FLINK_SQL_JOB            ListFlinkJobsRequestJobType
	FLINK_OPENSOURCE_SQL_JOB ListFlinkJobsRequestJobType
	FLINK_SQL_EDGE_JOB       ListFlinkJobsRequestJobType
	FLINK_JAR_JOB            ListFlinkJobsRequestJobType
}

func GetListFlinkJobsRequestJobTypeEnum() ListFlinkJobsRequestJobTypeEnum {
	return ListFlinkJobsRequestJobTypeEnum{
		FLINK_SQL_JOB: ListFlinkJobsRequestJobType{
			value: "flink_sql_job",
		},
		FLINK_OPENSOURCE_SQL_JOB: ListFlinkJobsRequestJobType{
			value: "flink_opensource_sql_job",
		},
		FLINK_SQL_EDGE_JOB: ListFlinkJobsRequestJobType{
			value: "flink_sql_edge_job",
		},
		FLINK_JAR_JOB: ListFlinkJobsRequestJobType{
			value: "flink_jar_job",
		},
	}
}

func (c ListFlinkJobsRequestJobType) Value() string {
	return c.value
}

func (c ListFlinkJobsRequestJobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFlinkJobsRequestJobType) UnmarshalJSON(b []byte) error {
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

type ListFlinkJobsRequestStatus struct {
	value string
}

type ListFlinkJobsRequestStatusEnum struct {
	JOB_INIT                 ListFlinkJobsRequestStatus
	JOB_SUBMITTING           ListFlinkJobsRequestStatus
	JOB_SUBMIT_FAIL          ListFlinkJobsRequestStatus
	JOB_RUNNING              ListFlinkJobsRequestStatus
	JOB_RUNNING_EXCEPTION    ListFlinkJobsRequestStatus
	JOB_DOWNLOADING          ListFlinkJobsRequestStatus
	JOB_IDLE                 ListFlinkJobsRequestStatus
	JOB_CANCELING            ListFlinkJobsRequestStatus
	JOB_CANCEL_SUCCESS       ListFlinkJobsRequestStatus
	JOB_CANCEL_FAIL          ListFlinkJobsRequestStatus
	JOB_SAVEPOINTING         ListFlinkJobsRequestStatus
	JOB_ARREARAGE_STOPPED    ListFlinkJobsRequestStatus
	JOB_ARREARAGE_RECOVERING ListFlinkJobsRequestStatus
	JOB_FINISH               ListFlinkJobsRequestStatus
}

func GetListFlinkJobsRequestStatusEnum() ListFlinkJobsRequestStatusEnum {
	return ListFlinkJobsRequestStatusEnum{
		JOB_INIT: ListFlinkJobsRequestStatus{
			value: "job_init",
		},
		JOB_SUBMITTING: ListFlinkJobsRequestStatus{
			value: "job_submitting",
		},
		JOB_SUBMIT_FAIL: ListFlinkJobsRequestStatus{
			value: "job_submit_fail",
		},
		JOB_RUNNING: ListFlinkJobsRequestStatus{
			value: "job_running",
		},
		JOB_RUNNING_EXCEPTION: ListFlinkJobsRequestStatus{
			value: "job_running_exception",
		},
		JOB_DOWNLOADING: ListFlinkJobsRequestStatus{
			value: "job_downloading",
		},
		JOB_IDLE: ListFlinkJobsRequestStatus{
			value: "job_idle",
		},
		JOB_CANCELING: ListFlinkJobsRequestStatus{
			value: "job_canceling",
		},
		JOB_CANCEL_SUCCESS: ListFlinkJobsRequestStatus{
			value: "job_cancel_success",
		},
		JOB_CANCEL_FAIL: ListFlinkJobsRequestStatus{
			value: "job_cancel_fail",
		},
		JOB_SAVEPOINTING: ListFlinkJobsRequestStatus{
			value: "job_savepointing",
		},
		JOB_ARREARAGE_STOPPED: ListFlinkJobsRequestStatus{
			value: "job_arrearage_stopped",
		},
		JOB_ARREARAGE_RECOVERING: ListFlinkJobsRequestStatus{
			value: "job_arrearage_recovering",
		},
		JOB_FINISH: ListFlinkJobsRequestStatus{
			value: "job_finish",
		},
	}
}

func (c ListFlinkJobsRequestStatus) Value() string {
	return c.value
}

func (c ListFlinkJobsRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFlinkJobsRequestStatus) UnmarshalJSON(b []byte) error {
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
