package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateResourcePackagesOrderRequest Request Object
type CreateResourcePackagesOrderRequest struct {
	Body *CreateResourcePackageOrderReq `json:"body,omitempty"`
}

func (o CreateResourcePackagesOrderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourcePackagesOrderRequest struct{}"
	}

	return strings.Join([]string{"CreateResourcePackagesOrderRequest", string(data)}, " ")
}
