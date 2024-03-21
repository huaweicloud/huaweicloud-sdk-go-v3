package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RenewDataProfileRequest Request Object
type RenewDataProfileRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)
	Workspace string `json:"workspace"`

	Body *DataProfileRo `json:"body,omitempty"`
}

func (o RenewDataProfileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RenewDataProfileRequest struct{}"
	}

	return strings.Join([]string{"RenewDataProfileRequest", string(data)}, " ")
}
