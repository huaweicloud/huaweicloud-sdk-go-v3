package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProjectBucketRsp 项目桶
type ProjectBucketRsp struct {

	// 项目编号
	EihealthProjectId *string `json:"eihealth_project_id,omitempty"`

	// 项目名称
	EihealthProjectName *string `json:"eihealth_project_name,omitempty"`

	// 桶类型(real:项目桶,quote:引用桶)
	Type *string `json:"type,omitempty"`

	// 是否引用桶根路径
	QuoteRoot *bool `json:"quote_root,omitempty"`
}

func (o ProjectBucketRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectBucketRsp struct{}"
	}

	return strings.Join([]string{"ProjectBucketRsp", string(data)}, " ")
}
