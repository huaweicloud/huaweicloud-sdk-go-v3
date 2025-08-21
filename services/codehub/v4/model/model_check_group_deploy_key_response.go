package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckGroupDeployKeyResponse Response Object
type CheckGroupDeployKeyResponse struct {

	// **参数解释：** 部署秘钥在上层代码组或项目是否配置。 **取值范围：** - true，上层代码组或项目已配置该密钥。 - false，上层代码组或项目未配置该密钥。
	Exists         *bool `json:"exists,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o CheckGroupDeployKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckGroupDeployKeyResponse struct{}"
	}

	return strings.Join([]string{"CheckGroupDeployKeyResponse", string(data)}, " ")
}
