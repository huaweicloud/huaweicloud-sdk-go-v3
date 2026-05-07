package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateHostResourceResponse Response Object
type BatchUpdateHostResourceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchUpdateHostResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateHostResourceResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateHostResourceResponse", string(data)}, " ")
}
