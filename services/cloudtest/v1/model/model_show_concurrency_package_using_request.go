package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConcurrencyPackageUsingRequest Request Object
type ShowConcurrencyPackageUsingRequest struct {

	// 服务id
	ServiceId string `json:"service_id"`

	// test_type
	TestType *string `json:"test_type,omitempty"`
}

func (o ShowConcurrencyPackageUsingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConcurrencyPackageUsingRequest struct{}"
	}

	return strings.Join([]string{"ShowConcurrencyPackageUsingRequest", string(data)}, " ")
}
