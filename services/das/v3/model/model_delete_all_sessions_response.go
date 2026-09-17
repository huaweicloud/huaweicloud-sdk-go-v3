package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAllSessionsResponse Response Object
type DeleteAllSessionsResponse struct {

	// 操作是否成功
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o DeleteAllSessionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAllSessionsResponse struct{}"
	}

	return strings.Join([]string{"DeleteAllSessionsResponse", string(data)}, " ")
}
