package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateQueueRequestBody 创建队列的请求参数。
type CreateQueueRequestBody struct {

	// 新建的队列名称，名称只能包含数字、英文字母和下划线，但不能是纯数字，且不能以下划线开头。长度限制：1~128个字符。\\n说明： \\n队列名称不区分大小写，系统会自动转换为小写。
	QueueName string `json:"queue_name"`

	// 队列的类型,。有如下两种类型： sql general 如果不指定，默认为sql。
	QueueType *string `json:"queue_type,omitempty"`

	// 队列的描述信息。
	Description *string `json:"description,omitempty"`

	// 队列的实际CU。
	CuCount int32 `json:"cu_count"`

	// 队列的收费模式。只能设置为“1”，表示按照CU时收费。
	ChargingMode *int32 `json:"charging_mode,omitempty"`

	// 企业项目ID，“0”表示default，即默认的企业项目。 说明： 开通了企业管理服务的用户可设置该参数绑定指定的项目。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 队列计算资源的cpu架构。
	Platform *string `json:"platform,omitempty"`

	// 队列资源模式。支持以下两种类型：0：共享资源模式1：专属资源模式
	ResourceMode *int32 `json:"resource_mode,omitempty"`

	// 创建队列的标签信息，目前包括队列是否跨AZ的标签信息（Json字符串），且只支持值为“2”，即创建一个计算资源分布在2个可用区的双AZ队列
	Labels *[]interface{} `json:"labels,omitempty"`

	// 队列特性。支持以下两种类型：basic：基础型ai：AI增强型（仅SQL的x86_64专属队列支持选择）默认值为“basic”。
	Feature *string `json:"feature,omitempty"`

	// 标签
	Tags *[]Tag `json:"tags,omitempty"`

	// 新建的弹性资源池名称，名称只能包含数字、小写英文字母和下划线，但不能是纯数字，且不能以下划线开头。长度限制：1~128个字符。
	ElasticResourcePoolName *string `json:"elastic_resource_pool_name,omitempty"`
}

func (o CreateQueueRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateQueueRequestBody struct{}"
	}

	return strings.Join([]string{"CreateQueueRequestBody", string(data)}, " ")
}
