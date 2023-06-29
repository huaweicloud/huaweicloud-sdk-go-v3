package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectInstanceTagsResponse Response Object
type ListProjectInstanceTagsResponse struct {

	// 项目下所有实例绑定的标签列表
	Tags           *[]TmsKeyValues `json:"tags,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListProjectInstanceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectInstanceTagsResponse struct{}"
	}

	return strings.Join([]string{"ListProjectInstanceTagsResponse", string(data)}, " ")
}
