package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 扩容实例磁盘时必填。 所需扩容到的磁盘容量大小。
type OpenGaussEnlargeVolume struct {

	// GaussDB磁盘大小要求（分片数*40GB）的倍数；取值范围：（分片数*40GB）~（分片数*16TB）
	Size int32 `json:"size"`
}

func (o OpenGaussEnlargeVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenGaussEnlargeVolume struct{}"
	}

	return strings.Join([]string{"OpenGaussEnlargeVolume", string(data)}, " ")
}
