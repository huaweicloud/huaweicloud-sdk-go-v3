package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAttributeResponse Response Object
type CreateAttributeResponse struct {

	// 自定义属性标识
	Id             *int64 `json:"id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateAttributeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAttributeResponse struct{}"
	}

	return strings.Join([]string{"CreateAttributeResponse", string(data)}, " ")
}
