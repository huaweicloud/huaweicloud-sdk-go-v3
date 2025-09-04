package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPackageSpecCountDownRequest Request Object
type ShowPackageSpecCountDownRequest struct {
	Body *CountdownRequestBody `json:"body,omitempty"`
}

func (o ShowPackageSpecCountDownRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPackageSpecCountDownRequest struct{}"
	}

	return strings.Join([]string{"ShowPackageSpecCountDownRequest", string(data)}, " ")
}
