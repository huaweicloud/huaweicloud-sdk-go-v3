package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type EnableLtsLogsRequest struct {
}

func (o EnableLtsLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableLtsLogsRequest struct{}"
	}

	return strings.Join([]string{"EnableLtsLogsRequest", string(data)}, " ")
}
