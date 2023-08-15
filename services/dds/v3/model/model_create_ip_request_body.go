package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateIpRequestBody struct {

	// 待打开IP开关的对象类型。 - 扩容shard组时，取值为“shard”。 - 扩容config组时，取值为“config”。
	Type string `json:"type"`

	// Shard组ID 注意：   1. 第一次添加Shard/Config IP时，该参数不传。   2. 对于已经添加过Shard IP的实例，需要传入该参数为新扩容的Shard组添加IP。
	TargetId *string `json:"target_id,omitempty"`

	// 打开集群开关设置的密码。 注意：该密码暂不支持修改，请谨慎操作。
	Password string `json:"password"`
}

func (o CreateIpRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIpRequestBody struct{}"
	}

	return strings.Join([]string{"CreateIpRequestBody", string(data)}, " ")
}
