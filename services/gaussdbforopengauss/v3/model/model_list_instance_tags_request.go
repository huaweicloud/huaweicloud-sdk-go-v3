package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceTagsRequest Request Object
type ListInstanceTagsRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ListInstanceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceTagsRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceTagsRequest", string(data)}, " ")
}
