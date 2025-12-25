package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceTagResponse Response Object
type ListResourceTagResponse struct {

	// 标签列表
	Tags           *[]WorkspaceTag `json:"tags,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListResourceTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceTagResponse struct{}"
	}

	return strings.Join([]string{"ListResourceTagResponse", string(data)}, " ")
}
