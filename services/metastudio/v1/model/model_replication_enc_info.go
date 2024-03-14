package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReplicationEncInfo 加密信息
type ReplicationEncInfo struct {

	// 加密算法
	Algorithm *string `json:"algorithm,omitempty"`

	// 秘钥id
	KeyId *string `json:"key_id,omitempty"`

	// 初始化向量
	Iv *string `json:"iv,omitempty"`
}

func (o ReplicationEncInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReplicationEncInfo struct{}"
	}

	return strings.Join([]string{"ReplicationEncInfo", string(data)}, " ")
}
