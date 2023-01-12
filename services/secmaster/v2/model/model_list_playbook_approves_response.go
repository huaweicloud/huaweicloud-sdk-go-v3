package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListPlaybookApprovesResponse struct {

	// Error code
	Code *string `json:"code,omitempty"`

	// Error message
	Message *string `json:"message,omitempty"`

	Data *[]ApproveOpinionDetail `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListPlaybookApprovesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPlaybookApprovesResponse struct{}"
	}

	return strings.Join([]string{"ListPlaybookApprovesResponse", string(data)}, " ")
}
