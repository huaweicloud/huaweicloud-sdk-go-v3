package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetJobBatchNameResponse Response Object
type SetJobBatchNameResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SetJobBatchNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetJobBatchNameResponse struct{}"
	}

	return strings.Join([]string{"SetJobBatchNameResponse", string(data)}, " ")
}
