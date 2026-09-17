package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddResourceToIteratorRequest Request Object
type AddResourceToIteratorRequest struct {

	// 迭代uri
	IteratorUri string `json:"iterator_uri"`

	// 是否异步返回, 默认false， 超过500时，前端传true
	IsAsync *bool `json:"is_async,omitempty"`

	Body *AddResourceInfo `json:"body,omitempty"`
}

func (o AddResourceToIteratorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddResourceToIteratorRequest struct{}"
	}

	return strings.Join([]string{"AddResourceToIteratorRequest", string(data)}, " ")
}
