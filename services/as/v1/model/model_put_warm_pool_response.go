package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PutWarmPoolResponse Response Object
type PutWarmPoolResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o PutWarmPoolResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutWarmPoolResponse struct{}"
	}

	return strings.Join([]string{"PutWarmPoolResponse", string(data)}, " ")
}
