package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApproversResponse Response Object
type ListApproversResponse struct {
	Data           *ListApproversResultData `json:"data,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListApproversResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApproversResponse struct{}"
	}

	return strings.Join([]string{"ListApproversResponse", string(data)}, " ")
}
