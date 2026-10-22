package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteV2TenantImageResponse Response Object
type DeleteV2TenantImageResponse struct {

	// **参数解释**：镜像ID。 **取值范围**：长度为32~36的英文字符、数字和中划线的组合。
	ImageId        *string `json:"image_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteV2TenantImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteV2TenantImageResponse struct{}"
	}

	return strings.Join([]string{"DeleteV2TenantImageResponse", string(data)}, " ")
}
