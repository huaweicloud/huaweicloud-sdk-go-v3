package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountInstancesByTagRequest Request Object
type CountInstancesByTagRequest struct {
	Body *ListCbhByTagsRequestBody `json:"body,omitempty"`
}

func (o CountInstancesByTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountInstancesByTagRequest struct{}"
	}

	return strings.Join([]string{"CountInstancesByTagRequest", string(data)}, " ")
}
