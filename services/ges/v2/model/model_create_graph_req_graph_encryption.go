package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGraphReqGraphEncryption 图实例是否加密，默认不加密。
type CreateGraphReqGraphEncryption struct {

	// 是否启动加密特性。取值为“true”或者“false”。默认为“false”。
	Enable *bool `json:"enable,omitempty"`

	// 与建图对应的project下，华为云数据加密服务创建的用户主密钥ID。
	MasterKeyId *string `json:"master_key_id,omitempty"`
}

func (o CreateGraphReqGraphEncryption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGraphReqGraphEncryption struct{}"
	}

	return strings.Join([]string{"CreateGraphReqGraphEncryption", string(data)}, " ")
}
