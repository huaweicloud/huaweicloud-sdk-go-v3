package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceResourceTagsResponse Response Object
type ListInstanceResourceTagsResponse struct {

	// 资源标签列表
	Tags           *[]ResourceTag `json:"tags,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListInstanceResourceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceResourceTagsResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceResourceTagsResponse", string(data)}, " ")
}
