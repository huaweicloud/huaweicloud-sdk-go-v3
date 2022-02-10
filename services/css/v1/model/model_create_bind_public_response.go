package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateBindPublicResponse struct {
	// 操作行为。

	Action         *string `json:"action,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateBindPublicResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBindPublicResponse struct{}"
	}

	return strings.Join([]string{"CreateBindPublicResponse", string(data)}, " ")
}
