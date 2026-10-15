package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRaySessionResponse Response Object
type CreateRaySessionResponse struct {

	// **参数解释**：会话ID。 **取值范围**：长度为1~36的英文、数字和中划线的组合。
	SessionId *string `json:"session_id,omitempty"`

	// **参数解释**：pod-agency委托Secret名称，用于挂载联邦身份凭证。非AIHub联邦认证场景下返回空字符串。 **约束限制**：不涉及。 **取值范围**：长度不超过256个字符，为空时返回空字符串。 **默认取值**：空字符串。
	SecretName *string `json:"secret_name,omitempty"`

	// **参数解释**：Ray集群GRPC连接端口号。 **约束限制**：不涉及。 **取值范围**：长度不超过5个字符，为空时返回空字符串。 **默认取值**：空字符串。
	RayGrpcServerPort *string `json:"ray_grpc_server_port,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateRaySessionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRaySessionResponse struct{}"
	}

	return strings.Join([]string{"CreateRaySessionResponse", string(data)}, " ")
}
