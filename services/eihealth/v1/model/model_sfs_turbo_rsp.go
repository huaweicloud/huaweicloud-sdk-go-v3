package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SfsTurboRsp 最终租户sfs-turbo资源详情。
type SfsTurboRsp struct {

	// sfs-turbo资源ID。
	Id *string `json:"id,omitempty"`

	// sfs-turbo资源名称。
	Name *string `json:"name,omitempty"`

	// sfs-turbo资源状态。
	Status *string `json:"status,omitempty"`

	// sfs-turbo资源类型。
	Type *string `json:"type,omitempty"`

	// sfs-turbo资源容量，单位GB。
	Space *string `json:"space,omitempty"`
}

func (o SfsTurboRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SfsTurboRsp struct{}"
	}

	return strings.Join([]string{"SfsTurboRsp", string(data)}, " ")
}
