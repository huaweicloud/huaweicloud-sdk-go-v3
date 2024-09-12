package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStoreRequest Request Object
type ListStoreRequest struct {
	Body *ListStoreRequestBody `bson:"body,omitempty"`
}

func (o ListStoreRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStoreRequest struct{}"
	}

	return strings.Join([]string{"ListStoreRequest", string(data)}, " ")
}
