package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportDlqMessageResponse Response Object
type ExportDlqMessageResponse struct {
	Body           *[]Message `json:"body,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ExportDlqMessageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportDlqMessageResponse struct{}"
	}

	return strings.Join([]string{"ExportDlqMessageResponse", string(data)}, " ")
}
