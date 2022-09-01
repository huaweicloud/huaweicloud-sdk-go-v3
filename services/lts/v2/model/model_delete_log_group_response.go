package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteLogGroupResponse struct {
	Body           *string `json:"body,omitempty" xml:"body"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteLogGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLogGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteLogGroupResponse", string(data)}, " ")
}
