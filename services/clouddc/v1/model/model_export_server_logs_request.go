package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportServerLogsRequest Request Object
type ExportServerLogsRequest struct {

	// imetal server id
	Id string `json:"id"`
}

func (o ExportServerLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportServerLogsRequest struct{}"
	}

	return strings.Join([]string{"ExportServerLogsRequest", string(data)}, " ")
}
