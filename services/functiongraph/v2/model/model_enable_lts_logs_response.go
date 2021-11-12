package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type EnableLtsLogsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o EnableLtsLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableLtsLogsResponse struct{}"
	}

	return strings.Join([]string{"EnableLtsLogsResponse", string(data)}, " ")
}
