package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type JobExecution struct {

	// 作业类型: - MapReduce - SparkSubmit - SparkPython：该类型作业将转换为SparkSubmit类型提交，MRS控制台界面的作业类型展示为SparkSubmit，通过接口查询作业列表信息时作业类型请选择SparkSubmit。 - HiveScript - HiveSql - DistCp，导入、导出数据。 - SparkScript - SparkSql - Flink
	JobType string `json:"job_type"`

	// 作业名称，只能由字母、数字、中划线和下划线组成，并且长度为1～64个字符。 说明： 不同作业的名称允许相同，但不建议设置相同。
	JobName string `json:"job_name"`

	// 程序执行的关键参数，该参数由用户程序内的函数指定，MRS只负责参数的传入。 最多为150000字符，不能包含;|&>'<$!\\\"\\特殊字符，可为空。 说明： - 若输入带有敏感信息（如登录密码）的参数可能在作业详情展示和日志打印中存在暴露的风险，请谨慎操作。 - 提交HiveScript或HiveSql类型的作业时如需以“obs://”开头格式访问存储在OBS上的文件，请在Hive服务配置页面搜索参数“core.site.customized.configs”，新增OBS的endpoint配置项，参数为“fs.obs.endpoint”，值请输入OBS对应的endpoint，具体请参考[终端节点](https://support.huaweicloud.com/api-mrs/mrs_02_0003.html)。
	Arguments *[]string `json:"arguments,omitempty"`

	// 程序系统参数。 最多为2048字符，不能包含><|'`&!\\特殊字符，可为空。
	Properties map[string]string `json:"properties,omitempty"`
}

func (o JobExecution) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobExecution struct{}"
	}

	return strings.Join([]string{"JobExecution", string(data)}, " ")
}
