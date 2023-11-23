package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInstancesSessionResponse Response Object
type DeleteInstancesSessionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteInstancesSessionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstancesSessionResponse struct{}"
	}

	return strings.Join([]string{"DeleteInstancesSessionResponse", string(data)}, " ")
}
