package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteApiToInstanceResponse Response Object
type ExecuteApiToInstanceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExecuteApiToInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteApiToInstanceResponse struct{}"
	}

	return strings.Join([]string{"ExecuteApiToInstanceResponse", string(data)}, " ")
}
