package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListStatusStatisticResponse struct {
	Body           *[]UserStatusStatistic `json:"body,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListStatusStatisticResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStatusStatisticResponse struct{}"
	}

	return strings.Join([]string{"ListStatusStatisticResponse", string(data)}, " ")
}
