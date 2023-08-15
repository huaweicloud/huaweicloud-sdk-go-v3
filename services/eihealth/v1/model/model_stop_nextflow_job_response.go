package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopNextflowJobResponse Response Object
type StopNextflowJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StopNextflowJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopNextflowJobResponse struct{}"
	}

	return strings.Join([]string{"StopNextflowJobResponse", string(data)}, " ")
}
