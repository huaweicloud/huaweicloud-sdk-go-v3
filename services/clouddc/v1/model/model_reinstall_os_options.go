package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReinstallOsOptions 重装OS请求参数，不会擦除数据盘
type ReinstallOsOptions struct {

	// 镜像ID，非必填，不传默认使用当前镜像ID
	ImageId *string `json:"image_id,omitempty"`

	// 设置实例的管理员账户初始登录密码，其中，Linux管理员账户为root，Windows管理员账户为Administrator。
	Password string `json:"password"`

	// **参数解释**： 实例id 列表 **约束限制**： 实例id不超过10条
	InstanceIdSet *[]string `json:"instance_id_set,omitempty"`
}

func (o ReinstallOsOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReinstallOsOptions struct{}"
	}

	return strings.Join([]string{"ReinstallOsOptions", string(data)}, " ")
}
