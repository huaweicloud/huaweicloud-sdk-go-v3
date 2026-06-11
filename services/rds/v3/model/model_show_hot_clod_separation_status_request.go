package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHotClodSeparationStatusRequest Request Object
type ShowHotClodSeparationStatusRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 语言。默认en-us。
	XLanguage *string `json:"X-Language,omitempty"`
}

func (o ShowHotClodSeparationStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHotClodSeparationStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowHotClodSeparationStatusRequest", string(data)}, " ")
}
