package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RetentionLog struct {

	// 创建时间
	CreatedAt string `json:"created_at" xml:"created_at"`

	// ID
	Id int32 `json:"id" xml:"id"`

	// 组织名
	Namespace string `json:"namespace" xml:"namespace"`

	// 镜像仓库名
	Repo string `json:"repo" xml:"repo"`

	// 老化规则ID
	RetentionId int32 `json:"retention_id" xml:"retention_id"`

	// 规则
	RuleType string `json:"rule_type" xml:"rule_type"`

	// 镜像版本
	Tag string `json:"tag" xml:"tag"`
}

func (o RetentionLog) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetentionLog struct{}"
	}

	return strings.Join([]string{"RetentionLog", string(data)}, " ")
}
