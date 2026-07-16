package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteImageGroupResponse Response Object
type DeleteImageGroupResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o DeleteImageGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteImageGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteImageGroupResponse", string(data)}, " ")
}
