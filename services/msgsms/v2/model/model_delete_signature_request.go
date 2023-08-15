package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSignatureRequest Request Object
type DeleteSignatureRequest struct {

	// 签名ID
	Id string `json:"id"`
}

func (o DeleteSignatureRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSignatureRequest struct{}"
	}

	return strings.Join([]string{"DeleteSignatureRequest", string(data)}, " ")
}
