package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStacksRequest Request Object
type ListStacksRequest struct {

	// 技术栈标签。默认为空值，查询全部。 例如：Java,CPP,GO,Python;可查询多个标签
	Tags *string `json:"tags,omitempty"`
}

func (o ListStacksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStacksRequest struct{}"
	}

	return strings.Join([]string{"ListStacksRequest", string(data)}, " ")
}
