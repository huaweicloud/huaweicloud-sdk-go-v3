package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ExecuteUpdateImageNameResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExecuteUpdateImageNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteUpdateImageNameResponse struct{}"
	}

	return strings.Join([]string{"ExecuteUpdateImageNameResponse", string(data)}, " ")
}
