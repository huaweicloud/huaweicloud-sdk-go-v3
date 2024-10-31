package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLimitTaskResponse Response Object
type DeleteLimitTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteLimitTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLimitTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteLimitTaskResponse", string(data)}, " ")
}
