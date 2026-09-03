package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteTestConnectionNewResponse Response Object
type ExecuteTestConnectionNewResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExecuteTestConnectionNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteTestConnectionNewResponse struct{}"
	}

	return strings.Join([]string{"ExecuteTestConnectionNewResponse", string(data)}, " ")
}
