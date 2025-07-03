package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeLargeVersionRequest Request Object
type UpgradeLargeVersionRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例id
	InstanceId string `json:"instance_id"`

	Body *CustomerUpgradeMajorVersionReq `json:"body,omitempty"`
}

func (o UpgradeLargeVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeLargeVersionRequest struct{}"
	}

	return strings.Join([]string{"UpgradeLargeVersionRequest", string(data)}, " ")
}
