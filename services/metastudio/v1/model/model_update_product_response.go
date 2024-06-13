package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateProductResponse Response Object
type UpdateProductResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateProductResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProductResponse struct{}"
	}

	return strings.Join([]string{"UpdateProductResponse", string(data)}, " ")
}
