package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssumeAgencyForPodIdentityRequestBody 使用ServiceAccount token获取对应的pod-identity关联相关委托临时凭据请求参数
type AssumeAgencyForPodIdentityRequestBody struct {

	// **参数解释：** pod-identity关联所绑定的ServiceAccount token。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 无
	Token string `json:"token"`
}

func (o AssumeAgencyForPodIdentityRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssumeAgencyForPodIdentityRequestBody struct{}"
	}

	return strings.Join([]string{"AssumeAgencyForPodIdentityRequestBody", string(data)}, " ")
}
