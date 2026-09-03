package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteLoginConnectionNewResponse Response Object
type ExecuteLoginConnectionNewResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExecuteLoginConnectionNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteLoginConnectionNewResponse struct{}"
	}

	return strings.Join([]string{"ExecuteLoginConnectionNewResponse", string(data)}, " ")
}
