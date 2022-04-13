package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DisassociateAppQuotaWithAppResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DisassociateAppQuotaWithAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateAppQuotaWithAppResponse struct{}"
	}

	return strings.Join([]string{"DisassociateAppQuotaWithAppResponse", string(data)}, " ")
}
