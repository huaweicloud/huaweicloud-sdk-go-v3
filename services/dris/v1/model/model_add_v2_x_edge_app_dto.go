package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddV2XEdgeAppDto AppModel传输对象
type AddV2XEdgeAppDto struct {

	// **参数说明**：用户自定义应用唯一ID，部署边缘应用前应先创建应用，方法参见：[创建应用](https://support.huaweicloud.com/api-v2x/v2x_04_0026.html)。  **取值范围**：只允许字母、数字、下划线（_）、连接符（-）、美元符号（$）的组合。
	EdgeAppId string `json:"edge_app_id"`

	// **参数说明**：应用版本，部署边缘应用版本前应先创建应用版本并发布，创建应用版本方法参见：[创建应用](https://support.huaweicloud.com/api-v2x/v2x_04_0038.html)，发布应用版本方法参见：[更新应用版本状态](https://support.huaweicloud.com/api-v2x/v2x_04_0105.html)。  **取值范围**：只允许小写字母、数字、连接符（-）、点（.）的组合且要以小写字母或数字开头和结尾。
	AppVersion string `json:"app_version"`
}

func (o AddV2XEdgeAppDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddV2XEdgeAppDto struct{}"
	}

	return strings.Join([]string{"AddV2XEdgeAppDto", string(data)}, " ")
}
