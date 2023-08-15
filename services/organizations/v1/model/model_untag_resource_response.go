package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UntagResourceResponse Response Object
type UntagResourceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UntagResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UntagResourceResponse struct{}"
	}

	return strings.Join([]string{"UntagResourceResponse", string(data)}, " ")
}
