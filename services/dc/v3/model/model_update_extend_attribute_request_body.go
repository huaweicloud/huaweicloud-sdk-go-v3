package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateExtendAttributeRequestBody 更新虚拟接口可靠性探测扩展信息的请求体
type UpdateExtendAttributeRequestBody struct {
	ExtendAttribute *VifExtendAttribute `json:"extend_attribute"`
}

func (o UpdateExtendAttributeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateExtendAttributeRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateExtendAttributeRequestBody", string(data)}, " ")
}
