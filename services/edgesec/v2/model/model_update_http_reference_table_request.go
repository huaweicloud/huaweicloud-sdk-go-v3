package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHttpReferenceTableRequest Request Object
type UpdateHttpReferenceTableRequest struct {

	// 引用表id
	TableId string `json:"table_id"`

	Body *UpdateHttpReferenceTableRequestBody `json:"body,omitempty"`
}

func (o UpdateHttpReferenceTableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHttpReferenceTableRequest struct{}"
	}

	return strings.Join([]string{"UpdateHttpReferenceTableRequest", string(data)}, " ")
}
