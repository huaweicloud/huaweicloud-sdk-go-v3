package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 批量作业详情
type JobContentInfo struct {

	// 批量作业对象类型，支持如下选项： - node：边缘节点 - node_group：边缘节点组 - deployment：边缘应用
	TargetType *string `json:"target_type,omitempty"`

	// 批量作业对象详情
	Targets *[]Target `json:"targets,omitempty"`

	// 批量作业内容，仅在批量应用部署和批量应用升级时需要填写，填入的内容为：使用json结构体编写的创建应用部署接口请求体deployment参数，并将其转换为字符串
	TaskData *string `json:"task_data,omitempty"`
}

func (o JobContentInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobContentInfo struct{}"
	}

	return strings.Join([]string{"JobContentInfo", string(data)}, " ")
}
