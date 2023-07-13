package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateStoredQueryRequest Request Object
type UpdateStoredQueryRequest struct {

	// 查询ID
	QueryId string `json:"query_id"`

	Body *StoredQueryRequestBody `json:"body,omitempty"`
}

func (o UpdateStoredQueryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStoredQueryRequest struct{}"
	}

	return strings.Join([]string{"UpdateStoredQueryRequest", string(data)}, " ")
}
