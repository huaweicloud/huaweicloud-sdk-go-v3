package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAccessControlLogsResponse Response Object
type ListAccessControlLogsResponse struct {
	Data           *HttpQueryCfwAccessControllerLogsResponseDtoData `json:"data,omitempty"`
	HttpStatusCode int                                              `json:"-"`
}

func (o ListAccessControlLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccessControlLogsResponse struct{}"
	}

	return strings.Join([]string{"ListAccessControlLogsResponse", string(data)}, " ")
}
