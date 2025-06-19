package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SparkJob 创建批处理作业的响应参数。
type SparkJob struct {

	// 参数解释:  Batch作业的id。 示例: 80ceaaff-3cfc-4162-a56f-70031ea4fa91 约束限制:  无 取值范围: 无 默认取值: 无
	Id *string `json:"id,omitempty"`

	// 参数解释:  Batch作业的状态 示例: starting 约束限制:  无 取值范围: starting：正在启动；running：正在执行任务；dead：session已退出；success：session停止成功；recovering：正在恢复 默认取值: 无
	State *string `json:"state,omitempty"`

	// 参数解释:  批处理作业的后台app id 示例: batch-session-1f49b734-757a-419c-9519-7754520cb03c:31309 约束限制:  无 取值范围: 无 默认取值: 无
	AppId *string `json:"appId,omitempty"`

	// 参数解释:  显示当前Batch作业的最后10条记录 约束限制:  无 取值范围: 无 默认取值: 无
	Log *[]string `json:"log,omitempty"`

	// 参数解释:   计算资源类型，目前可接受参数A, B, C。如果不指定，则按最小类型创建。 示例: A 约束限制:  无 取值范围: A：物理资源：8核32G内存，driverCores：2；executorCores：1；driverMemory：7G；executorMemory：4G；numExecutor：6。 B：16核64G内存,2,2,7G,8G,7。 C：32核128G内存,4,2,15G,8G,14。 默认取值: 无
	ScType *string `json:"sc_type,omitempty"`

	// 参数解释:  会话所在队列 示例: test 约束限制:  无 取值范围: 无 默认取值: 无
	ClusterName *string `json:"cluster_name,omitempty"`

	// 参数解释:  Batch的创建时间。是单位为“毫秒”的时间戳 示例: 1747169165821 约束限制:  无 取值范围: 大于等于0的整数 默认取值: 无
	CreateTime *int64 `json:"create_time,omitempty"`

	// 参数解释:  创建时用户指定的批处理名称，不能超过128个字符 示例: test_pyFiles 约束限制:  无 取值范围: 无 默认取值: 无
	Name *string `json:"name,omitempty"`

	// 参数解释:  批处理作业所属用户 示例: ei_dlics_d00352431 约束限制:  无 取值范围: 无 默认取值: 无
	Owner *string `json:"owner,omitempty"`

	// 参数解释:  批处理作业所属代理用户（资源租户） 示例: tenant1 约束限制:  无 取值范围: 无 默认取值: 无
	ProxyUser *string `json:"proxyUser,omitempty"`

	// 参数解释:   批处理作业类型，只支持spark类型参数 示例: spark 约束限制:  无 取值范围: spark 默认取值: 无
	Kind *string `json:"kind,omitempty"`

	// 参数解释:   用于指定队列，填写已创建DLI的队列名 示例: gen_native 约束限制:  无 取值范围: 无 默认取值: 无
	Queue *string `json:"queue,omitempty"`

	// 参数解释:   自定义镜像。格式为：组织名/镜像名:镜像版本 示例: ceshi/spark_general-x86_64:3.3.1-2.3.8.1020240906885119450448000 约束限制:  格式为: 组织名/镜像名:镜像版本的字符串 取值范围: 无 默认取值: 无
	Image *string `json:"image,omitempty"`

	// 参数解释:   请求参数详情 示例: {\\\"jars\\\":[\\\"spark-examples_2.11-2.4.5-h0.cbu.dli.300.20240725.r1.jar\\\"],\\\"pyFiles\\\":[],\\\"files\\\":[],\\\"modelFiles\\\":[],\\\"resources\\\":[],\\\"modules\\\":[],\\\"groups\\\":[],\\\"archives\\\":[],\\\"queue\\\":\\\"gen0218\\\",\\\"name\\\":\\\"\\\",\\\"conf\\\":{},\\\"execution_agency_urn\\\":\\\"agency\\\",\\\"file\\\":\\\"obs://dli-wzy/package/job/spark/longrunning.py\\\",\\\"args\\\":[],\\\"className\\\":\\\"\\\",\\\"autoRecovery\\\":false,\\\"minRecoveryDelayTime\\\":10000,\\\"maxRetryTimes\\\":20,\\\"obs_bucket\\\":\\\"rain3\\\",\\\"image\\\":\\\"ceshi/spark_general-x86_64:3.3.1-2.3.8.1020240906885119450448000\\\",\\\"feature\\\":\\\"custom\\\",\\\"spark_version\\\":\\\"3.3.1\\\"} 约束限制:  符合Json格式的字符串 取值范围: 无 默认取值: 无
	ReqBody *string `json:"req_body,omitempty"`

	// 参数解释:   更新时间 示例: 1739867996988 约束限制:  无 取值范围: 大于等于0的整数 默认取值: 无
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 参数解释:   作业运行时长，单位毫秒 示例: 141079 约束限制:  无 取值范围: 大于等于0的整数 默认取值: 无
	Duration *int64 `json:"duration,omitempty"`
}

func (o SparkJob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SparkJob struct{}"
	}

	return strings.Join([]string{"SparkJob", string(data)}, " ")
}
