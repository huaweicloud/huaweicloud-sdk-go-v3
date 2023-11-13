package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCenterTaskResponse Response Object
type DeleteCenterTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteCenterTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCenterTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteCenterTaskResponse", string(data)}, " ")
}
