package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DisassociateSignatureKeyV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DisassociateSignatureKeyV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateSignatureKeyV2Response struct{}"
	}

	return strings.Join([]string{"DisassociateSignatureKeyV2Response", string(data)}, " ")
}
