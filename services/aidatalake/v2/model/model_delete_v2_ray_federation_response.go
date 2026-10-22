package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteV2RayFederationResponse Response Object
type DeleteV2RayFederationResponse struct {

	// **参数解释**：Ray联邦端点ID。 **取值范围**：长度为32~36的英文字符、数字和中划线的组合。
	RayFederationId *string `json:"ray_federation_id,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteV2RayFederationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteV2RayFederationResponse struct{}"
	}

	return strings.Join([]string{"DeleteV2RayFederationResponse", string(data)}, " ")
}
