package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableLtsLogsRequest Request Object
type EnableLtsLogsRequest struct {
}

func (o EnableLtsLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableLtsLogsRequest struct{}"
	}

	return strings.Join([]string{"EnableLtsLogsRequest", string(data)}, " ")
}
