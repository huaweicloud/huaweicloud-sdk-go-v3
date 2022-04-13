package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RegisterImageRequest struct {
	// 镜像ID。 image_id为用户调用创建镜像元数据接口所创建出来镜像的id，使用其他方式创建的镜像id会导致注册失败。 注册接口调用成功后，请根据镜像id查询镜像的状态。镜像状态变为active表示镜像注册成功，详情请参见查询镜像详情（OpenStack原生）。

	ImageId string `json:"image_id"`

	Body *RegisterImageRequestBody `json:"body,omitempty"`
}

func (o RegisterImageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterImageRequest struct{}"
	}

	return strings.Join([]string{"RegisterImageRequest", string(data)}, " ")
}
