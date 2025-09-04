package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIteratorsInfoRequest Request Object
type ListIteratorsInfoRequest struct {
	Body *IteratorVersionsQueryInfo `json:"body,omitempty"`
}

func (o ListIteratorsInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIteratorsInfoRequest struct{}"
	}

	return strings.Join([]string{"ListIteratorsInfoRequest", string(data)}, " ")
}
