package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Statement struct {

	// 授权项。指对资源的具体操作权限，不超过100个。 - 格式为：服务名:资源类型:操作，例：vpc:ports:create。 - 服务名为产品名称，例如ecs、evs和vpc等，服务名仅支持小写。 资源类型和操作没有大小写，要求支持通配符号*，无需罗列全部授权项。 - 当自定义策略为委托自定义策略时，该字段值为： \"Action\": [\"iam:agencies:assume\"]。
	Action *[]string `json:"action,omitempty"`

	// 资源。数组长度不超过10，每个字符串长度不超过128，规则如下： - 可填 * 的五段式：::::，例：\"obs:::bucket:*\"。 - region字段为*或用户可访问的region。service必须存在且resource属于对应service。 - 当该自定义策略为委托自定义策略时，该字段类型为Object，值为：\"Resource\": {\"uri\": [\"/iam/agencies/07805acaba800fdd4fbdc00b8f888c7c\"]}。
	Resource *[]string `json:"resource,omitempty"`
}

func (o Statement) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Statement struct{}"
	}

	return strings.Join([]string{"Statement", string(data)}, " ")
}
