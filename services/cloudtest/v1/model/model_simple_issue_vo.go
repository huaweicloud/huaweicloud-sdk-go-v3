package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SimpleIssueVo 需求
type SimpleIssueVo struct {

	// ID
	Id *string `json:"id,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 层级路径
	Path *string `json:"path,omitempty"`

	// 类型
	TrackerName *string `json:"tracker_name,omitempty"`
}

func (o SimpleIssueVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SimpleIssueVo struct{}"
	}

	return strings.Join([]string{"SimpleIssueVo", string(data)}, " ")
}
