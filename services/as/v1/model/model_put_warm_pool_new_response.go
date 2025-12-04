package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PutWarmPoolNewResponse Response Object
type PutWarmPoolNewResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o PutWarmPoolNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutWarmPoolNewResponse struct{}"
	}

	return strings.Join([]string{"PutWarmPoolNewResponse", string(data)}, " ")
}
