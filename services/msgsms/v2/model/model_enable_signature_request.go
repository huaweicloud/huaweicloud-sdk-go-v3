package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableSignatureRequest Request Object
type EnableSignatureRequest struct {

	// 签名ID
	Id string `json:"id"`
}

func (o EnableSignatureRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableSignatureRequest struct{}"
	}

	return strings.Join([]string{"EnableSignatureRequest", string(data)}, " ")
}
