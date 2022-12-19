package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AppModel传输对象
type ModifyV2XEdgeAppDto struct {

	// **参数说明**：应用版本，比如1.0.0，升级边缘应用版本时应确保应用版本状态为发布（PUBLISHED），查询应用版本状态方法参见：[查询应用版本详情](https://support.huaweicloud.com/api-v2x/v2x_04_0104.html)，更新应用版本状态方法参见：[更新应用版本状态](https://support.huaweicloud.com/api-v2x/v2x_04_0105.html)。  **取值范围**：只允许小写字母、数字、连接符（-）、点（.）的组合且要以小写字母或数字开头和结尾。
	AppVersion string `json:"app_version"`
}

func (o ModifyV2XEdgeAppDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyV2XEdgeAppDto struct{}"
	}

	return strings.Join([]string{"ModifyV2XEdgeAppDto", string(data)}, " ")
}
