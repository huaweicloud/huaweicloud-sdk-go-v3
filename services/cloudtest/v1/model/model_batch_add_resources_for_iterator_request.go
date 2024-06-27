package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAddResourcesForIteratorRequest Request Object
type BatchAddResourcesForIteratorRequest struct {

	// 迭代uri
	IteratorId string `json:"iterator_id"`

	Body *AddResourceInfo `json:"body,omitempty"`
}

func (o BatchAddResourcesForIteratorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddResourcesForIteratorRequest struct{}"
	}

	return strings.Join([]string{"BatchAddResourcesForIteratorRequest", string(data)}, " ")
}
