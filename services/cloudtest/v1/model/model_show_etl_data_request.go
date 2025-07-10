package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEtlDataRequest Request Object
type ShowEtlDataRequest struct {
	Body *EtlRequestBody `json:"body,omitempty"`
}

func (o ShowEtlDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEtlDataRequest struct{}"
	}

	return strings.Join([]string{"ShowEtlDataRequest", string(data)}, " ")
}
