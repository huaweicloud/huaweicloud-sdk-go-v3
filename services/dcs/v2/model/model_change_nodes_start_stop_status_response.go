package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeNodesStartStopStatusResponse Response Object
type ChangeNodesStartStopStatusResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeNodesStartStopStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeNodesStartStopStatusResponse struct{}"
	}

	return strings.Join([]string{"ChangeNodesStartStopStatusResponse", string(data)}, " ")
}
