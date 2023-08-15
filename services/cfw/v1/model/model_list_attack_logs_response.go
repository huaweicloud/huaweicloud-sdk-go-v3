package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAttackLogsResponse Response Object
type ListAttackLogsResponse struct {
	Data           *HttpQueryCfwAttackLogsResponseDtoData `json:"data,omitempty"`
	HttpStatusCode int                                    `json:"-"`
}

func (o ListAttackLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAttackLogsResponse struct{}"
	}

	return strings.Join([]string{"ListAttackLogsResponse", string(data)}, " ")
}
