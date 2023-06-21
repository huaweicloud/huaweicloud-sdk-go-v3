package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UnlockPortResponse struct {
	Data           *UnlockPortResponseModel `json:"data,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o UnlockPortResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnlockPortResponse struct{}"
	}

	return strings.Join([]string{"UnlockPortResponse", string(data)}, " ")
}
