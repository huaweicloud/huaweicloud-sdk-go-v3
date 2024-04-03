package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SparkJob 创建批处理作业的响应参数。
type SparkJob struct {

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

	// 请求参数详情
	ReqBody *string `json:"req_body,omitempty"`

	// 更新时间
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 作业运行时长，单位毫秒。
	Duration *int64 `json:"duration,omitempty"`
}

func (o SparkJob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SparkJob struct{}"
	}

	return strings.Join([]string{"SparkJob", string(data)}, " ")
}
