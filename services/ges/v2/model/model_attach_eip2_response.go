package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachEip2Response Response Object
type AttachEip2Response struct {
	Body           map[string]string `json:"body,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o AttachEip2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachEip2Response struct{}"
	}

	return strings.Join([]string{"AttachEip2Response", string(data)}, " ")
}
