package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFabricProjectTagsResponse Response Object
type ListFabricProjectTagsResponse struct {

	// 标签列表
	Tags *[]ResourceTagParam `json:"tags,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListFabricProjectTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFabricProjectTagsResponse struct{}"
	}

	return strings.Join([]string{"ListFabricProjectTagsResponse", string(data)}, " ")
}
