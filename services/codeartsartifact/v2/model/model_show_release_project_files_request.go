package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowReleaseProjectFilesRequest Request Object
type ShowReleaseProjectFilesRequest struct {

	// 文件名称，模糊搜索
	FileName string `json:"file_name"`

	// 每页显示的条目数量
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0
	Offset *int32 `json:"offset,omitempty"`

	// 华为云devcloud的项目id
	ProjectId string `json:"project_id"`
}

func (o ShowReleaseProjectFilesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReleaseProjectFilesRequest struct{}"
	}

	return strings.Join([]string{"ShowReleaseProjectFilesRequest", string(data)}, " ")
}
