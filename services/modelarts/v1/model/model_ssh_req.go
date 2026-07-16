package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SshReq SSH连接信息。
type SshReq struct {

	// SSH密钥对名称，可以在云服务器控制台（ECS）“密钥对”页面创建和查看。
	KeyPairNames *[]string `json:"key_pair_names,omitempty"`
}

func (o SshReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SshReq struct{}"
	}

	return strings.Join([]string{"SshReq", string(data)}, " ")
}
