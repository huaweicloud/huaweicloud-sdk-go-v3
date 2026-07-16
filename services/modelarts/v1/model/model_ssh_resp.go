package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SshResp SSH连接信息。
type SshResp struct {

	// SSH密钥对名称，可以在云服务器控制台（ECS）“密钥对”页面创建和查看。
	KeyPairNames []string `json:"key_pair_names"`

	// SSH连接地址信息。
	TaskUrls *[]TaskUrls `json:"task_urls,omitempty"`
}

func (o SshResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SshResp struct{}"
	}

	return strings.Join([]string{"SshResp", string(data)}, " ")
}
