package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLogtanksResponse Response Object
type ListLogtanksResponse struct {
	Logtanks *[]LogtankDetail `json:"logtanks,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListLogtanksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogtanksResponse struct{}"
	}

	return strings.Join([]string{"ListLogtanksResponse", string(data)}, " ")
}
