package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBatchInfoResponse Response Object
type ShowBatchInfoResponse struct {

	// Batch作业的id。
	Id *string `json:"id,omitempty"`

	// Batch作业的状态。包括： starting：正在启动；running：正在执行任务；dead：session已退出；success：session停止成功；recovering：正在恢复。
	State *string `json:"state,omitempty"`

	// 批处理作业的后台app id。
	AppId *string `json:"appId,omitempty"`

	// 显示当前Batch作业的最后10条记录。
	Log *[]string `json:"log,omitempty"`

	// 计算资源类型。用户自定义时返回CUSTOMIZED。
	ScType *string `json:"sc_type,omitempty"`

	// 会话所在队列。
	ClusterName *string `json:"cluster_name,omitempty"`

	// Batch的创建时间。是单位为“毫秒”的时间戳。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 创建时用户指定的批处理名称，不能超过128个字符。
	Name *string `json:"name,omitempty"`

	// 批处理作业所属用户
	Owner *string `json:"owner,omitempty"`

	// 批处理作业所属代理用户（资源租户）。
	ProxyUser *string `json:"proxyUser,omitempty"`

	// 批处理作业类型，只支持spark类型参数。
	Kind *string `json:"kind,omitempty"`

	// 用于指定队列，填写已创建DLI的队列名
	Queue *string `json:"queue,omitempty"`

	// 自定义镜像。格式为：组织名/镜像名:镜像版本。
	Image *string `json:"image,omitempty"`

	// 更新时间
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 作业特性。表示用户作业使用的Spark镜像类型。  basic：表示使用DLI提供的基础Spark镜像。 custom：表示使用用户自定义的Spark镜像。 ai：表示使用DLI提供的AI镜像。
	Feature *string `json:"feature,omitempty"`

	// 作业使用spark组件的版本号，在“feature”为“basic”或“ai”时填写，若不填写，则使用默认的spark组件版本号2.3.2。
	SparkVersion   *string `json:"spark_version,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowBatchInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBatchInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowBatchInfoResponse", string(data)}, " ")
}
