package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReinstallVolumeSpec 服务器重装云硬盘配置
type ReinstallVolumeSpec struct {

	// **参数解释**： 用户自定义镜像ID > 获取方式:在控制台的“服务列表”中，单击“计算 > 镜像服务 > 私有镜像”，单击镜像的名称，在服务器详情页的“基本信息”页签下找到“镜像ID”字段复制即可。  **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	ImageID *string `json:"imageID,omitempty"`

	// **参数解释**： 用户主密钥ID。默认为空时，表示云硬盘不加密。 [> 获取密钥ID的方法请参考：[查询密钥列表](https://support.huaweicloud.com/api-dew/ListKeys.html)](tag:hws) [> 获取密钥ID的方法请参考：[查询密钥列表](https://support.huaweicloud.com/intl/zh-cn/api-dew/ListKeys.html)](tag:hws_hk)  **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	CmkID *string `json:"cmkID,omitempty"`
}

func (o ReinstallVolumeSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReinstallVolumeSpec struct{}"
	}

	return strings.Join([]string{"ReinstallVolumeSpec", string(data)}, " ")
}
