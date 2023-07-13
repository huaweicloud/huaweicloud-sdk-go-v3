package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronCreateNetworkRequest Request Object
type NeutronCreateNetworkRequest struct {
	Body *NeutronCreateNetworkRequestBody `json:"body,omitempty"`
}

func (o NeutronCreateNetworkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronCreateNetworkRequest struct{}"
	}

	return strings.Join([]string{"NeutronCreateNetworkRequest", string(data)}, " ")
}
