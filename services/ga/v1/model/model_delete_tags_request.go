package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTagsRequest Request Object
type DeleteTagsRequest struct {

	// 资源类型。
	ResourceType *ResourceType `json:"resource_type"`

	// 资源ID。
	ResourceId string `json:"resource_id"`

	Body *DeleteTagsRequestBody `json:"body,omitempty"`
}

func (o DeleteTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTagsRequest struct{}"
	}

	return strings.Join([]string{"DeleteTagsRequest", string(data)}, " ")
}
