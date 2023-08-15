package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAppConfigV2Request Request Object
type CreateAppConfigV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 应用编号
	AppId string `json:"app_id"`

	Body *AppConfigCreateRequestV2 `json:"body,omitempty"`
}

func (o CreateAppConfigV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppConfigV2Request struct{}"
	}

	return strings.Join([]string{"CreateAppConfigV2Request", string(data)}, " ")
}
