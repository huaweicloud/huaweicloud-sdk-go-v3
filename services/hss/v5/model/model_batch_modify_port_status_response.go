package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchModifyPortStatusResponse Response Object
type BatchModifyPortStatusResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchModifyPortStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchModifyPortStatusResponse struct{}"
	}

	return strings.Join([]string{"BatchModifyPortStatusResponse", string(data)}, " ")
}
