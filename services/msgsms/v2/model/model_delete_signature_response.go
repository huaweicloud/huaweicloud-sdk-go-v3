package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSignatureResponse Response Object
type DeleteSignatureResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSignatureResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSignatureResponse struct{}"
	}

	return strings.Join([]string{"DeleteSignatureResponse", string(data)}, " ")
}
