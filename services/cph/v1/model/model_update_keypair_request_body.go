package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateKeypairRequestBody struct {

	// 待更改密钥对的云手机服务器信息。
	Servers []ServerKeypair `json:"servers"`
}

func (o UpdateKeypairRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKeypairRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateKeypairRequestBody", string(data)}, " ")
}
