package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteProductResponse Response Object
type DeleteProductResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteProductResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteProductResponse struct{}"
	}

	return strings.Join([]string{"DeleteProductResponse", string(data)}, " ")
}
