package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteApproverResponse Response Object
type DeleteApproverResponse struct {
	Data           *DeleteResultData `json:"data,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o DeleteApproverResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteApproverResponse struct{}"
	}

	return strings.Join([]string{"DeleteApproverResponse", string(data)}, " ")
}
