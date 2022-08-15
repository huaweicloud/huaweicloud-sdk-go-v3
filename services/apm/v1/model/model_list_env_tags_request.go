package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListEnvTagsRequest struct {

	// 应用id
	XBusinessId int64 `json:"x-business-id"`

	Body *TagParam `json:"body,omitempty"`
}

func (o ListEnvTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnvTagsRequest struct{}"
	}

	return strings.Join([]string{"ListEnvTagsRequest", string(data)}, " ")
}
