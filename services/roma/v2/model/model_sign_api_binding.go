package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SignApiBinding struct {
	// 签名密钥编号

	SignId string `json:"sign_id"`
	// API的发布记录编号

	PublishIds []string `json:"publish_ids"`
}

func (o SignApiBinding) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SignApiBinding struct{}"
	}

	return strings.Join([]string{"SignApiBinding", string(data)}, " ")
}
