package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 数据盘信息
type CreateDataImage struct {

	// 数据盘镜像名称。
	Name string `json:"name" xml:"name"`

	// 数据盘ID。
	VolumeId string `json:"volume_id" xml:"volume_id"`

	// 数据盘描述。
	Description *string `json:"description,omitempty" xml:"description"`

	// 数据盘镜像标签。
	Tags *[]string `json:"tags,omitempty" xml:"tags"`
}

func (o CreateDataImage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDataImage struct{}"
	}

	return strings.Join([]string{"CreateDataImage", string(data)}, " ")
}
