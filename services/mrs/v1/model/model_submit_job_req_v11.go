package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SubmitJobReqV11 struct {

	// 作业名称，只能由字母、数字、中划线和下划线组成，并且长度为1～64个字符。   说明： 不同作业的名称允许相同，但不建议设置相同。
	JobName string `json:"job_name" xml:"job_name"`

	// 集群ID。
	ClusterId string `json:"cluster_id" xml:"cluster_id"`

	// 执行程序Jar包或sql文件地址，需要满足如下要求：  - 最多为1023字符，不能包含;|&><'$特殊字符，且不可为空或全空格。 - 需要以“/”或“s3a://”开头。OBS路径不支持KMS加密的文件或程序。 - Spark Script需要以“.sql”结尾，MapReduce和Spark Jar需要以“.jar”结尾，sql和jar不区分大小写。  说明： 作业类型为MapReduce或Spark时，jar_path参数为必选。
	JarPath *string `json:"jar_path,omitempty" xml:"jar_path"`

	// 数据输入地址，必须以“/”或“s3a://”开头。请配置为正确的OBS路径，OBS路径不支持KMS加密的文件或程序。  最多为1023字符，不能包含;|&>'<$特殊字符，可为空。
	Input *string `json:"input,omitempty" xml:"input"`

	// 数据输出地址，必须以“/”或“s3a://”开头。请配置为正确的OBS路径，如果该路径不存在，系统会自动创建。  最多为1023字符，不能包含;|&>'<$特殊字符，可为空。
	Output *string `json:"output,omitempty" xml:"output"`

	// 作业日志存储地址，该日志信息记录作业运行状态。必须以“/”或“s3a://”开头，请配置为正确的OBS路径。  最多为1023字符，不能包含;|&>'<$特殊字符，可为空。
	JobLog *string `json:"job_log,omitempty" xml:"job_log"`

	// 作业类型码。  - 1：MapReduce - 2：Spark - 3：Hive Script - 4：HiveSQL（当前不支持） - 5：DistCp，导入、导出数据。 - 6：Spark Script - 7：Spark SQL，提交SQL语句（该接口当前不支持）
	JobType int32 `json:"job_type" xml:"job_type"`

	//   文件操作类型，包括： export：从HDFS导出数据至OBS import：从OBS导入数据至HDFS
	FileAction *string `json:"file_action,omitempty" xml:"file_action"`

	// 程序执行的关键参数，该参数由用户程序内的函数指定，MRS只负责参数的传入。 最多为150000字符，不能包含;|&>'<$!\\\"\\特殊字符，可为空。 说明： 用户输入带有敏感信息（如登录密码）的参数时，可通过在参数名前添加“@”的方式，为该参数值加密，以防止敏感信息被明文形式持久化。在查看作业信息时，敏感信息显示为“*”。 例如：username=admin @password=admin_123
	Arguments *string `json:"arguments,omitempty" xml:"arguments"`

	// Spark SQL语句，该语句需要进行Base64编码和解码，“ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/”为标准的编码表，MRS使用“ABCDEFGHILKJMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/”进行Base64编码。在编码后所得字符串首位任意加上一个字母，即得到Hql参数的值。后台自动进行解码得到Spark SQL语句。 使用样例： 1) 在Web界面输入Spark SQL语句“show tables;”。 2) 使用“ABCDEFGHILKJMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/”编码后得到字符串“c2hvdyB0YWLsZXM7”。 3) 在“c2hvdyB0YWLsZXM7”首位任意加上一字母，例如“gc2hvdyB0YWLsZXM7”，即Hql参数的值。 4) 后台自动进行解码得到Spark SQL语句“show tables;”。
	Hql *string `json:"hql,omitempty" xml:"hql"`

	// sql程序路径，仅Spark Script和Hive Script作业需要使用此参数。需要满足如下要求：  - 最多为1023字符，不能包含;|&><'$特殊字符，且不可为空或全空格。 - 需要以“/”或“s3a://”开头，OBS路径不支持KMS加密的文件或程序。 - 需要以“.sql”结尾，sql不区分大小写。
	HiveScriptPath *string `json:"hive_script_path,omitempty" xml:"hive_script_path"`
}

func (o SubmitJobReqV11) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubmitJobReqV11 struct{}"
	}

	return strings.Join([]string{"SubmitJobReqV11", string(data)}, " ")
}
