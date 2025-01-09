package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCsrRequest Request Object
type UpdateCsrRequest struct {

	// CSR的ID。
	Id string `json:"id"`

	Body *UpdateCsrRequestBody `json:"body,omitempty"`
}

func (o UpdateCsrRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCsrRequest struct{}"
	}

	return strings.Join([]string{"UpdateCsrRequest", string(data)}, " ")
}
