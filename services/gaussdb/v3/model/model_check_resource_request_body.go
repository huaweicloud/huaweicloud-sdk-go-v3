package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CheckResourceRequestBody struct {

	// 校验类型。 - createInstance：校验创建实例。 - createReadonlyNode：校验创建只读节点。 - resizeFlavor：校验规格变更。
	Action string `json:"action"`

	Resource *CheckResourceInfo `json:"resource"`
}

func (o CheckResourceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckResourceRequestBody struct{}"
	}

	return strings.Join([]string{"CheckResourceRequestBody", string(data)}, " ")
}
