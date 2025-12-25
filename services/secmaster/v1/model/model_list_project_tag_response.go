package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectTagResponse Response Object
type ListProjectTagResponse struct {

	// 标签列表
	Tags           *[]WorkspaceTag `json:"tags,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListProjectTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectTagResponse struct{}"
	}

	return strings.Join([]string{"ListProjectTagResponse", string(data)}, " ")
}
