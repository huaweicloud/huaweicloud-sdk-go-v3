package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteStackSetResponse Response Object
type DeleteStackSetResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteStackSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteStackSetResponse struct{}"
	}

	return strings.Join([]string{"DeleteStackSetResponse", string(data)}, " ")
}
