package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateIpRequestBody struct {

	// 待打开IP开关的对象类型。 - 扩容shard组时，取值为“shard”。 - 扩容config组时，取值为“config”。
	Type string `json:"type"`

	// 待打开IP开关的组ID。   - 对于shard组，取值为shard组ID。   - 对于config组，取值为config组ID。   - 如果为空，则打开该实例下同group类型的所有开关。 注意：   1. 第一次打开实例开关， 该参数需要传空。   2. 针对已开启开关的组， 开关不允许重复下发。
	TargetId *string `json:"target_id,omitempty"`

	// 打开集群开关设置的密码。  注意：该密码暂不支持修改，请谨慎操作。
	Password string `json:"password"`
}

func (o CreateIpRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIpRequestBody struct{}"
	}

	return strings.Join([]string{"CreateIpRequestBody", string(data)}, " ")
}
