package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListStacksByTagRequest struct {
	// 技术栈标签。默认为空值，查询全部。 例如：Java,CPP,GO,Python;可查询多个标签

	Tags *[]string `json:"tags,omitempty"`
}

func (o ListStacksByTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStacksByTagRequest struct{}"
	}

	return strings.Join([]string{"ListStacksByTagRequest", string(data)}, " ")
}
