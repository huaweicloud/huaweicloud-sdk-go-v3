package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetInstanceDataDumpResponse Response Object
type SetInstanceDataDumpResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SetInstanceDataDumpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetInstanceDataDumpResponse struct{}"
	}

	return strings.Join([]string{"SetInstanceDataDumpResponse", string(data)}, " ")
}
