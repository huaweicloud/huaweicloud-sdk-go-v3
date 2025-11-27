package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowExportTaskResponse Response Object
type ShowExportTaskResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowExportTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowExportTaskResponse struct{}"
	}

	return strings.Join([]string{"ShowExportTaskResponse", string(data)}, " ")
}
