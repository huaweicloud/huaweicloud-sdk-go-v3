package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTagRequest **参数解释**：给资源添加标签请求体。
type CreateTagRequest struct {

	// **参数解释**：待添加的标签列表。
	Tags []Tag `json:"tags"`
}

func (o CreateTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTagRequest struct{}"
	}

	return strings.Join([]string{"CreateTagRequest", string(data)}, " ")
}
