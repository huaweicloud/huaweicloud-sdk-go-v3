package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachEip2Request Request Object
type AttachEip2Request struct {

	// 图ID。
	GraphId string `json:"graph_id"`

	Body *AttachEipReq `json:"body,omitempty"`
}

func (o AttachEip2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachEip2Request struct{}"
	}

	return strings.Join([]string{"AttachEip2Request", string(data)}, " ")
}
