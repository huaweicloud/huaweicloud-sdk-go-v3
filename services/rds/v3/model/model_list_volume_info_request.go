package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVolumeInfoRequest Request Object
type ListVolumeInfoRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`
}

func (o ListVolumeInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVolumeInfoRequest struct{}"
	}

	return strings.Join([]string{"ListVolumeInfoRequest", string(data)}, " ")
}
