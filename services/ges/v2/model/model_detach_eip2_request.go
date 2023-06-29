package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetachEip2Request Request Object
type DetachEip2Request struct {

	// 图ID。
	GraphId string `json:"graph_id"`

	Body *DetachEipReq `json:"body,omitempty"`
}

func (o DetachEip2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachEip2Request struct{}"
	}

	return strings.Join([]string{"DetachEip2Request", string(data)}, " ")
}
