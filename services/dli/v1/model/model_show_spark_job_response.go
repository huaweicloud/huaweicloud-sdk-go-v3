package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSparkJobResponse Response Object
type ShowSparkJobResponse struct {

	// 参数解释:  Batch作业的id 示例: 80ceaaff-3cfc-4162-a56f-70031ea4fa91 约束限制:  无 取值范围: 无 默认取值: 无
	Id *string `json:"id,omitempty"`

	// 参数解释:  Batch作业的状态 示例: starting 约束限制:  无 取值范围: starting：正在启动 running：正在执行任务 dead：session已退出 success：session停止成功 recovering：正在恢复 默认取值: 无
	State *string `json:"state,omitempty"`

	// 参数解释:  批处理作业的后台app id 示例: batch-session-83ce2d53-8ae4-4afa-a0e2-e71a7aaa015c:30663 约束限制:  无 取值范围: 无 默认取值: 无
	AppId *string `json:"appId,omitempty"`

	// 参数解释:  显示当前Batch作业的最后10条记录 约束限制:  无 取值范围: 无 默认取值: 无
	Log *[]string `json:"log,omitempty"`

	// 参数解释:   计算资源类型。用户自定义时返回CUSTOMIZED 示例: CUSTOMIZED 约束限制:  无 取值范围: 无 默认取值: 无
	ScType *string `json:"sc_type,omitempty"`

	// 参数解释:   批处理作业所在队列 示例: test 约束限制:  无 取值范围: 无 默认取值: 无
	ClusterName *string `json:"cluster_name,omitempty"`

	// 参数解释:   Batch的创建时间。是单位为“毫秒”的时间戳 示例: 1747169026784 约束限制:  无 取值范围: 大于等于0的整数 默认取值: 无
	CreateTime *int64 `json:"create_time,omitempty"`

	// 参数解释:   创建时用户指定的批处理名称，不能超过128个字符 示例: testsparksdhd_smn 约束限制:  不超过128个字符的字符串 取值范围: 无 默认取值: 无
	Name *string `json:"name,omitempty"`

	// 参数解释:   批处理作业所属用户 示例: ei_dlics_d00352431 约束限制:  无 取值范围: 无 默认取值: 无
	Owner *string `json:"owner,omitempty"`

	// 参数解释:   批处理作业所属代理用户（资源租户） 示例: tenant1 约束限制:  无 取值范围: 无 默认取值: 无
	ProxyUser *string `json:"proxyUser,omitempty"`

	// 参数解释:   批处理作业类型，只支持spark类型参数 示例: spark 约束限制:  无 取值范围: spark 默认取值: 无
	Kind *string `json:"kind,omitempty"`

	// 参数解释:   用于指定队列，填写已创建DLI的队列名 示例: genzy_520 约束限制:  无 取值范围: 无 默认取值: 无
	Queue *string `json:"queue,omitempty"`

	// 参数解释:   自定义镜像。格式为：组织名/镜像名:镜像版本 示例: dli/spark:2.4.8 约束限制:  格式为：组织名/镜像名:镜像版本的字符串 取值范围: 无 默认取值: 无
	Image *string `json:"image,omitempty"`

	// 参数解释:   请求参数详情 示例: {\\\"jars\\\":[],\\\"pyFiles\\\":[],\\\"files\\\":[],\\\"modelFiles\\\":[],\\\"resources\\\":[],\\\"modules\\\":[],\\\"groups\\\":[],\\\"archives\\\":[],\\\"queue\\\":\\\"general_0509\\\",\\\"name\\\":\\\"testsparkjar43535\\\",\\\"conf\\\":{},\\\"execution_agency_urn\\\":\\\"agency\\\",\\\"file\\\":\\\"obs://qinzhuo/spark-examples_2.11-2.1.0.luxor.jar\\\",\\\"args\\\":[],\\\"className\\\":\\\"org.apache.spark.examples.SparkPi\\\",\\\"autoRecovery\\\":false,\\\"minRecoveryDelayTime\\\":10000,\\\"maxRetryTimes\\\":20,\\\"obs_bucket\\\":\\\"rain3\\\",\\\"feature\\\":\\\"basic\\\",\\\"spark_version\\\":\\\"3.3.1\\\"} 约束限制:  符合Json格式的字符串 取值范围: 无 默认取值: 无
	ReqBody *string `json:"req_body,omitempty"`

	// 参数解释:   更新时间 示例: 1747362066342 约束限制:  无 取值范围: 大于等于0的整数 默认取值: 无
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 参数解释:  作业特性。表示用户作业使用的Spark镜像类型。 示例: basic 约束限制:  无 取值范围: basic：表示使用DLI提供的基础Spark镜像 custom：表示使用用户自定义的Spark镜像 ai：表示使用DLI提供的AI镜像 默认取值: 无
	Feature *string `json:"feature,omitempty"`

	// 参数解释:  作业使用spark组件的版本号，在“feature”为“basic”或“ai”时填写，若不填写，则使用默认的spark组件版本号2.3.2 示例: 2.3.2 约束限制:  无 取值范围: 无 默认取值: 2.3.2
	SparkVersion   *string `json:"spark_version,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSparkJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSparkJobResponse struct{}"
	}

	return strings.Join([]string{"ShowSparkJobResponse", string(data)}, " ")
}
