package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type FreezeCertResponse struct {
	// 操作结果

	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o FreezeCertResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FreezeCertResponse struct{}"
	}

	return strings.Join([]string{"FreezeCertResponse", string(data)}, " ")
}
