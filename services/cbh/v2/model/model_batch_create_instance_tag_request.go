package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateInstanceTagRequest Request Object
type BatchCreateInstanceTagRequest struct {

	// 资源ID。(list接口获取)
	ResourceId string `json:"resource_id"`

	Body *CbsGetResourceIdTags `json:"body,omitempty"`
}

func (o BatchCreateInstanceTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateInstanceTagRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateInstanceTagRequest", string(data)}, " ")
}
