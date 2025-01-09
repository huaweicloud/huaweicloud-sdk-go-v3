package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Sfs3Storage 存储信息。
type Sfs3Storage struct {

	// 文件系统名称。
	Name *string `json:"name,omitempty"`

	// 挂载地址。
	Location *string `json:"location,omitempty"`

	// 存储使用量(Byte)。
	Usage *string `json:"usage,omitempty"`

	// 创建时间。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`
}

func (o Sfs3Storage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Sfs3Storage struct{}"
	}

	return strings.Join([]string{"Sfs3Storage", string(data)}, " ")
}
