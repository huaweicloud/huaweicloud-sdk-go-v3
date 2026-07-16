package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImageSaveJob struct {

	// 镜像名称，长度限制512个字符，支持小写字母、数字、中划线、下划线和点。
	Name *string `json:"name,omitempty"`

	// 镜像所属组织，可以在SWR控制台“组织管理”创建和查看。
	Namespace *string `json:"namespace,omitempty"`

	// 镜像tag，长度限制64个字符， 支持大小写字母、数字、中划线、下划线和点。
	Tag *string `json:"tag,omitempty"`

	// 该镜像所对应的描述信息，长度限制512个字符。
	Description *string `json:"description,omitempty"`
}

func (o ImageSaveJob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageSaveJob struct{}"
	}

	return strings.Join([]string{"ImageSaveJob", string(data)}, " ")
}
