package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowImageUploadCommandRequest Request Object
type ShowImageUploadCommandRequest struct {

	// **参数解释**： 镜像仓地址 **约束限制**： 网址/IP:端口。如：myharbor.com。 **取值范围**： 字符长度1-256位  **默认取值**： 不涉及
	RegistryAddr string `json:"registry_addr"`

	// **参数解释**： 镜像仓项目,用来指定扫描组件要上传到的镜像仓目录。 **约束限制**： 不涉及 **取值范围**： 字符长度1-128位  **默认取值**： 不涉及
	Namespace string `json:"namespace"`

	// **参数解释**： 用于登录镜像仓的用户名。 **约束限制**： 不涉及 **取值范围**： 字符长度1-128位  **默认取值**： 不涉及
	Username string `json:"username"`

	// **参数解释**： 用于登录镜像仓的密码。仅用于生成命令，HSS服务不会存储您的镜像仓密码。 **约束限制**： 不涉及 **取值范围**： 字符长度1-128位  **默认取值**： 不涉及
	Password string `json:"password"`
}

func (o ShowImageUploadCommandRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageUploadCommandRequest struct{}"
	}

	return strings.Join([]string{"ShowImageUploadCommandRequest", string(data)}, " ")
}
