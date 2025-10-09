package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetCustomizeTagsResponse Response Object
type GetCustomizeTagsResponse struct {

	// **参数解释**： 自定义标签 **约束限制**： 不涉及
	Tags           *[]CustomizeResourceTag `json:"tags,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o GetCustomizeTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetCustomizeTagsResponse struct{}"
	}

	return strings.Join([]string{"GetCustomizeTagsResponse", string(data)}, " ")
}
