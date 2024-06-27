package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIteratorsRequest Request Object
type ListIteratorsRequest struct {
	Body *IteratorVersionsQueryInfo `json:"body,omitempty"`
}

func (o ListIteratorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIteratorsRequest struct{}"
	}

	return strings.Join([]string{"ListIteratorsRequest", string(data)}, " ")
}
