package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReplayErrorClassification 回放异常SQL分类数据项
type ReplayErrorClassification struct {

	// 目标库标识
	TargetName *string `json:"target_name,omitempty"`

	// 异常SQL分类
	ErrorType string `json:"error_type"`

	// 异常SQL数量
	ErrorCnt *string `json:"error_cnt,omitempty"`

	// 异常SQL模板数量
	ErrorTemplateCnt *string `json:"error_template_cnt,omitempty"`
}

func (o ReplayErrorClassification) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReplayErrorClassification struct{}"
	}

	return strings.Join([]string{"ReplayErrorClassification", string(data)}, " ")
}
