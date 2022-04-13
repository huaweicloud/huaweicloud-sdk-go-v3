package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateTracingResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateTracingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTracingResponse struct{}"
	}

	return strings.Join([]string{"UpdateTracingResponse", string(data)}, " ")
}
