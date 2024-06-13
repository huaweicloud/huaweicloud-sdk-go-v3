package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListReportsRequest Request Object
type ListReportsRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	// 每页显示的条目数量,最大支持200条
	PageSize int64 `json:"page_size"`

	// 页数，page_no大于等于1
	Offset int64 `json:"offset"`

	// 版本id
	VersionId string `json:"version_id"`

	// 报表类型 1：首页用例库， 2：质量报告
	Type int32 `json:"type"`
}

func (o ListReportsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListReportsRequest struct{}"
	}

	return strings.Join([]string{"ListReportsRequest", string(data)}, " ")
}
