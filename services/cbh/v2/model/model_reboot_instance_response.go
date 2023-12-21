package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RebootInstanceResponse Response Object
type RebootInstanceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RebootInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RebootInstanceResponse struct{}"
	}

	return strings.Join([]string{"RebootInstanceResponse", string(data)}, " ")
}
