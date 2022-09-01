package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateAndExecuteJobResponse struct {

	// 作业执行对象是否由作业模板生成。
	Templated *bool `json:"templated,omitempty" xml:"templated"`

	// 作业创建时间，十位时间戳。
	CreatedAt *int64 `json:"created_at,omitempty" xml:"created_at"`

	// 作业更新时间，十位时间戳。
	UpdatedAt *int64 `json:"updated_at,omitempty" xml:"updated_at"`

	// 作业ID。
	Id *string `json:"id,omitempty" xml:"id"`

	// 项目编号。获取方法，请参见[获取项目ID](https://support.huaweicloud.com/api-mrs/mrs_02_0011.html)。
	TenantId *string `json:"tenant_id,omitempty" xml:"tenant_id"`

	// 作业应用ID。
	JobId *string `json:"job_id,omitempty" xml:"job_id"`

	// 作业名称，只能由字母、数字、中划线和下划线组成，并且长度为1～64个字符。  说明： 不同作业的名称允许相同，但不建议设置相同。
	JobName *string `json:"job_name,omitempty" xml:"job_name"`

	//   数据输入ID。
	InputId *string `json:"input_id,omitempty" xml:"input_id"`

	// 数据输出ID。
	OutputId *string `json:"output_id,omitempty" xml:"output_id"`

	// 作业执行开始时间，十位时间戳。
	StartTime *int64 `json:"start_time,omitempty" xml:"start_time"`

	// 作业执行结束时间，十位时间戳。
	EndTime *int64 `json:"end_time,omitempty" xml:"end_time"`

	// 集群ID。
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id"`

	// Oozie工作流ID。
	EngineJobId *string `json:"engine_job_id,omitempty" xml:"engine_job_id"`

	// 运行结果返回码。
	ReturnCode *string `json:"return_code,omitempty" xml:"return_code"`

	// 是否公开。 当前版本不支持该功能。
	IsPublic *bool `json:"is_public,omitempty" xml:"is_public"`

	// 是否受保护。 当前版本不支持该功能。
	IsProtected *bool `json:"is_protected,omitempty" xml:"is_protected"`

	// 作业执行组ID。
	GroupId *string `json:"group_id,omitempty" xml:"group_id"`

	// 执行程序Jar包或sql文件地址，需要满足如下要求：   - 最多为1023字符，不能包含;|&><'$特殊字符，且不可为空或全空格。   - 需要以“/”或“s3a://”开头。OBS路径不支持KMS加密的文件或程序。  - Spark Script需要以“.sql”结尾，MapReduce和Spark Jar需要以“.jar”结尾，sql和jar不区分大小写。
	JarPath *string `json:"jar_path,omitempty" xml:"jar_path"`

	// 数据输入地址，必须以“/”或“s3a://”开头。请配置为正确的OBS路径，OBS路径不支持KMS加密的文件或程序。  最多为1023字符，不能包含;|&>'<$特殊字符，可为空。
	Input *string `json:"input,omitempty" xml:"input"`

	// 数据输出地址，必须以“/”或“s3a://”开头。请配置为正确的OBS路径，如果该路径不存在，系统会自动创建。  最多为1023字符，不能包含;|&>'<$特殊字符，可为空。
	Output *string `json:"output,omitempty" xml:"output"`

	// 作业日志存储地址，该日志信息记录作业运行状态。必须以“/”或“s3a://”开头，请配置为正确的OBS路径。  最多为1023字符，不能包含;|&>'<$特殊字符，可为空。
	JobLog *string `json:"job_log,omitempty" xml:"job_log"`

	// 作业类型码。 - 1：MapReduce - 2：Spark - 3：Hive Script - 4：HiveSQL（当前不支持） - 5：DistCp，导入、导出数据。 - 6：Spark Script - 7：Spark SQL，提交SQL语句，（该接口当前不支持）  说明： 只有包含Spark和Hive组件的集群才能新增Spark和Hive类型的作业。
	JobType *int32 `json:"job_type,omitempty" xml:"job_type"`

	//   文件操作类型，包括： - export：从HDFS导出数据至OBS - import：从OBS导入数据至HDFS
	FileAction *string `json:"file_action,omitempty" xml:"file_action"`

	// 程序执行的关键参数，该参数由用户程序内的函数指定，MRS只负责参数的传入。 最多为150000字符，不能包含;|&>'<$!\\\"\\特殊字符，可为空。 说明： 用户输入带有敏感信息（如登录密码）的参数时，可通过在参数名前添加“@”的方式，为该参数值加密，以防止敏感信息被明文形式持久化。在查看作业信息时，敏感信息显示为“*”。 例如：username=admin @password=admin_123
	Arguments *string `json:"arguments,omitempty" xml:"arguments"`

	// Hive&Spark Sql语句
	Hql *string `json:"hql,omitempty" xml:"hql"`

	// 作业状态码。  - -1：Terminated - 1：Starting - 2：Running - 3：Completed - 4：Abnormal - 5：Error
	JobState *int32 `json:"job_state,omitempty" xml:"job_state"`

	// 作业最终状态码。  - 0：未完成 - 1：执行错误，终止执行 - 2：执行完成并且成功 - 3：已取消
	JobFinalStatus *int32 `json:"job_final_status,omitempty" xml:"job_final_status"`

	// sql程序路径，仅Spark Script和Hive Script作业需要使用此参数。需要满足如下要求：  - 最多为1023字符，不能包含;|&><'$特殊字符，且不可为空或全空格。 - 需要以“/”或“s3a://”开头，OBS路径不支持KMS加密的文件或程序。 - 需要以“.sql”结尾，sql不区分大小写。
	HiveScriptPath *string `json:"hive_script_path,omitempty" xml:"hive_script_path"`

	// 创建作业的用户ID。  为兼容历史版本，保留此参数。
	CreateBy *string `json:"create_by,omitempty" xml:"create_by"`

	// 当前已完成的步骤数。  为兼容历史版本，保留此参数。
	FinishedStep *int32 `json:"finished_step,omitempty" xml:"finished_step"`

	// 作业主ID。  为兼容历史版本，保留此参数。
	JobMainId *string `json:"job_main_id,omitempty" xml:"job_main_id"`

	// 作业步骤ID。  为兼容历史版本，保留此参数。
	JobStepId *string `json:"job_step_id,omitempty" xml:"job_step_id"`

	// 延迟时间，十位时间戳。  为兼容历史版本，保留此参数。
	PostponeAt *int64 `json:"postpone_at,omitempty" xml:"postpone_at"`

	// 作业步骤名。  为兼容历史版本，保留此参数。
	StepName *string `json:"step_name,omitempty" xml:"step_name"`

	// 步骤数量  为兼容历史版本，保留此参数。
	StepNum *int32 `json:"step_num,omitempty" xml:"step_num"`

	//   任务数量。 为兼容历史版本，保留此参数。
	TaskNum *int32 `json:"task_num,omitempty" xml:"task_num"`

	// 更新作业的用户ID。
	UpdateBy *string `json:"update_by,omitempty" xml:"update_by"`

	// 令牌，当前版本不支持。
	Credentials *string `json:"credentials,omitempty" xml:"credentials"`

	// 创建作业的用户ID。  历史版本兼容，不再使用。
	UserId *string `json:"user_id,omitempty" xml:"user_id"`

	// 键值对集合，用于保存作业运行配置。
	JobConfigs map[string]interface{} `json:"job_configs,omitempty" xml:"job_configs"`

	// 认证信息，当前版本不支持。
	Extra map[string]interface{} `json:"extra,omitempty" xml:"extra"`

	// 数据源URL。
	DataSourceUrls map[string]interface{} `json:"data_source_urls,omitempty" xml:"data_source_urls"`

	// 键值对集合，包含oozie返回的作业运行信息。
	Info           map[string]interface{} `json:"info,omitempty" xml:"info"`
	HttpStatusCode int                    `json:"-"`
}

func (o CreateAndExecuteJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAndExecuteJobResponse struct{}"
	}

	return strings.Join([]string{"CreateAndExecuteJobResponse", string(data)}, " ")
}
