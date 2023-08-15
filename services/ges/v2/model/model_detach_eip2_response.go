package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetachEip2Response Response Object
type DetachEip2Response struct {
	Body           map[string]string `json:"body,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o DetachEip2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachEip2Response struct{}"
	}

	return strings.Join([]string{"DetachEip2Response", string(data)}, " ")
}
