package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApiBatchPublish struct {

	// 需要发布或下线的API ID列表
	Apis *[]string `json:"apis,omitempty"`

	// 环境ID
	EnvId string `json:"env_id"`

	// 对本次发布的描述信息  字符长度不超过255 > 中文字符必须为UTF-8或者unicode编码。
	Remark *string `json:"remark,omitempty"`
}

func (o ApiBatchPublish) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiBatchPublish struct{}"
	}

	return strings.Join([]string{"ApiBatchPublish", string(data)}, " ")
}
