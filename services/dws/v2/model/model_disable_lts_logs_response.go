package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableLtsLogsResponse Response Object
type DisableLtsLogsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DisableLtsLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableLtsLogsResponse struct{}"
	}

	return strings.Join([]string{"DisableLtsLogsResponse", string(data)}, " ")
}
