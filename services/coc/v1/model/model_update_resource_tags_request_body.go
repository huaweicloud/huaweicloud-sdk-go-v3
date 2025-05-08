package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateResourceTagsRequestBody struct {
	Tags *[]ResourceTag `json:"tags,omitempty"`
}

func (o UpdateResourceTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResourceTagsRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateResourceTagsRequestBody", string(data)}, " ")
}
