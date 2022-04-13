package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchDeleteIterationsV4Response struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteIterationsV4Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteIterationsV4Response struct{}"
	}

	return strings.Join([]string{"BatchDeleteIterationsV4Response", string(data)}, " ")
}
