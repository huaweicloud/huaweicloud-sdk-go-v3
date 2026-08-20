package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteConnectionNewResponse Response Object
type BatchDeleteConnectionNewResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteConnectionNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteConnectionNewResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteConnectionNewResponse", string(data)}, " ")
}
