package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateStoredQueryRequest Request Object
type CreateStoredQueryRequest struct {
	Body *StoredQueryRequestBody `json:"body,omitempty"`
}

func (o CreateStoredQueryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStoredQueryRequest struct{}"
	}

	return strings.Join([]string{"CreateStoredQueryRequest", string(data)}, " ")
}
