package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTaskRecordResponse Response Object
type DeleteTaskRecordResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteTaskRecordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTaskRecordResponse struct{}"
	}

	return strings.Join([]string{"DeleteTaskRecordResponse", string(data)}, " ")
}
