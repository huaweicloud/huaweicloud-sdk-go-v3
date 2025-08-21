package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLogsResponse Response Object
type ListLogsResponse struct {
	Data           *ListLogsRespData `json:"data,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogsResponse struct{}"
	}

	return strings.Join([]string{"ListLogsResponse", string(data)}, " ")
}
