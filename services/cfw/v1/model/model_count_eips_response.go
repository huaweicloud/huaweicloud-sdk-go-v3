package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CountEipsResponse struct {

	// 防护对象ID
	ObjectId *string `json:"object_id,omitempty"`

	// EIP总数
	EipTotal *int32 `json:"eip_total,omitempty"`

	// EIP防护数
	EipProtected   *int32 `json:"eip_protected,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CountEipsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountEipsResponse struct{}"
	}

	return strings.Join([]string{"CountEipsResponse", string(data)}, " ")
}
