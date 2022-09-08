package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddJobs struct {

	// 作业类型码。 - 1：MapReduce - 2：Spark - 3：Hive Script - 4：HiveSQL（当前不支持） - 5：DistCp，导入、导出数据，（当前不支持）。 - 6：Spark Script - 7：Spark SQL，提交SQL语句，（当前不支持）。
	JobType int32 `json:"job_type"`

	// 作业名称。 只能由字母、数字、中划线和下划线组成，并且长度为1～64个字符。  说明： 不同作业的名称允许相同，但不建议设置相同。
	JobName string `json:"job_name"`

	// 执行程序Jar包或sql文件地址，需要满足如下要求： - 最多为1023字符，不能包含;|&>,<'$特殊字符，且不可为空或全空格。 - 文件可存储于HDFS或者OBS中，不同的文件系统对应的路径存在差异。    - OBS：以“obs://”开头。不支持KMS加密的文件或程序。    - HDFS：以“/”开头。 - Spark Script需要以“.sql”结尾，MapReduce和Spark Jar需要以“.jar”结尾，sql和jar不区分大小写。
	JarPath *string `json:"jar_path,omitempty"`

	// 程序执行的关键参数，该参数由用户程序内的函数指定，MRS只负责参数的传入。 最多为150000字符，不能包含;|&>'<$特殊字符，可为空。
	Arguments *string `json:"arguments,omitempty"`

	// 数据输入地址。 文件可存储于HDFS或者OBS中，不同的文件系统对应的路径存在差异。 - OBS：以“obs://”开头。不支持KMS加密的文件或程序。 - HDFS：以“/”开头。  最多为1023字符，不能包含;|&>'<$特殊字符，可为空。
	Input *string `json:"input,omitempty"`

	// 数据输出地址。 文件可存储于HDFS或者OBS中，不同的文件系统对应的路径存在差异。 - OBS：以“obs://”开头。 - HDFS：以“/”开头。  如果该路径不存在，系统会自动创建。 最多为1023字符，不能包含;|&>'<$特殊字符，可为空。
	Output *string `json:"output,omitempty"`

	// 作业日志存储地址，该日志信息记录作业运行状态。 文件可存储于HDFS或者OBS中，不同的文件系统对应的路径存在差异。 - OBS：以“obs://”开头。 - HDFS：以“/”开头。  最多为1023字符，不能包含;|&>'<$特殊字符，可为空。
	JobLog *string `json:"job_log,omitempty"`

	// sql程序路径，仅Spark Script和Hive Script作业需要使用此参数。需要满足如下要求： - 最多为1023字符，不能包含;|&><'$特殊字符，且不可为空或全空格。 - 文件可存储于HDFS或者OBS中，不同的文件系统对应的路径存在差异。     - OBS：以“obs://”开头。不支持KMS加密的文件或程序。     - HDFS：以“/”开头。 - 需要以“.sql”结尾，sql不区分大小写。
	HiveScriptPath *string `json:"hive_script_path,omitempty"`

	// HQL脚本语句。
	Hql *string `json:"hql,omitempty"`

	// 作业执行完成后，是否删除集群。  - true：是  - false：否
	ShutdownCluster *bool `json:"shutdown_cluster,omitempty"`

	// - true：创建集群同时提交作业  - false：单独提交作业  此处应设置为true。
	SubmitJobOnceClusterRun bool `json:"submit_job_once_cluster_run"`

	// 数据导入导出。 - import - export
	FileAction *string `json:"file_action,omitempty"`
}

func (o AddJobs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddJobs struct{}"
	}

	return strings.Join([]string{"AddJobs", string(data)}, " ")
}
