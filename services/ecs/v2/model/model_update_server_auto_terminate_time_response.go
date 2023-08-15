package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateServerAutoTerminateTimeResponse Response Object
type UpdateServerAutoTerminateTimeResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateServerAutoTerminateTimeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServerAutoTerminateTimeResponse struct{}"
	}

	return strings.Join([]string{"UpdateServerAutoTerminateTimeResponse", string(data)}, " ")
}
