package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBriefRecordResponseBodyResult 结果
type ListBriefRecordResponseBodyResult struct {

	// 需要查询的项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 简要构建信息列表
	BriefBuildRecordDtos *[]BriefRecordItems `json:"brief_build_record_dtos,omitempty"`
}

func (o ListBriefRecordResponseBodyResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBriefRecordResponseBodyResult struct{}"
	}

	return strings.Join([]string{"ListBriefRecordResponseBodyResult", string(data)}, " ")
}
