package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSharesByTagRequest Request Object
type ListSharesByTagRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type"`

	Body *ListSharesByTagRequestBody `json:"body,omitempty"`
}

func (o ListSharesByTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSharesByTagRequest struct{}"
	}

	return strings.Join([]string{"ListSharesByTagRequest", string(data)}, " ")
}
