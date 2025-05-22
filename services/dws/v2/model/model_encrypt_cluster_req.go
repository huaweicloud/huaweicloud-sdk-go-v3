package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EncryptClusterReq **参数解释**： 转加密集群请求信息。 **取值范围**： 不涉及。
type EncryptClusterReq struct {
	Encrypt *EncryptCluster `json:"encrypt,omitempty"`
}

func (o EncryptClusterReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EncryptClusterReq struct{}"
	}

	return strings.Join([]string{"EncryptClusterReq", string(data)}, " ")
}
