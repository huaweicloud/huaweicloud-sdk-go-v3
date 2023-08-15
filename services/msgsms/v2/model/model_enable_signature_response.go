package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableSignatureResponse Response Object
type EnableSignatureResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o EnableSignatureResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableSignatureResponse struct{}"
	}

	return strings.Join([]string{"EnableSignatureResponse", string(data)}, " ")
}
