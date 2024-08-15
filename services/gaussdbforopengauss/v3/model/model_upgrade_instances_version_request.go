package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeInstancesVersionRequest Request Object
type UpgradeInstancesVersionRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	Body *GaussDbUpgradeInstancesVersionRequest `json:"body,omitempty"`
}

func (o UpgradeInstancesVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeInstancesVersionRequest struct{}"
	}

	return strings.Join([]string{"UpgradeInstancesVersionRequest", string(data)}, " ")
}
