package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAttendeesResponse Response Object
type DeleteAttendeesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAttendeesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAttendeesResponse struct{}"
	}

	return strings.Join([]string{"DeleteAttendeesResponse", string(data)}, " ")
}
