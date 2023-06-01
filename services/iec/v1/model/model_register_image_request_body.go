package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type RegisterImageRequestBody struct {

	// 注册到边缘云上的公有云IMS的私有镜像id。
	ImageId *string `json:"image_id,omitempty"`

	// 原私有镜像所在公有云的region。
	RegionId *string `json:"region_id,omitempty"`
}

func (o RegisterImageRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterImageRequestBody struct{}"
	}

	return strings.Join([]string{"RegisterImageRequestBody", string(data)}, " ")
}
