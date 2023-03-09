package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListPublisherRequest struct {

	// 发布商名称
	Name *string `json:"name,omitempty"`
}

func (o ListPublisherRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPublisherRequest struct{}"
	}

	return strings.Join([]string{"ListPublisherRequest", string(data)}, " ")
}
