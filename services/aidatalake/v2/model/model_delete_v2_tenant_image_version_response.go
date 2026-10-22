package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteV2TenantImageVersionResponse Response Object
type DeleteV2TenantImageVersionResponse struct {

	// **参数解释**：镜像版本ID。 **取值范围**：长度为32~36的英文字符、数字和中划线的组合。
	VersionId *string `json:"version_id,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteV2TenantImageVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteV2TenantImageVersionResponse struct{}"
	}

	return strings.Join([]string{"DeleteV2TenantImageVersionResponse", string(data)}, " ")
}
