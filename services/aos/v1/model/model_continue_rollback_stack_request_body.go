package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ContinueRollbackStackRequestBody struct {

	// 资源栈（stack）的唯一ID。  此ID由资源编排服务在生成资源栈的时候生成，为UUID。  由于资源栈名仅仅在同一时间下唯一，即用户允许先生成一个叫HelloWorld的资源栈，删除，再重新创建一个同名资源栈。  对于团队并行开发，用户可能希望确保，当前我操作的资源栈就是我认为的那个，而不是其他队友删除后创建的同名资源栈。因此，使用ID就可以做到强匹配。  资源编排服务保证每次创建的资源栈所对应的ID都不相同，更新不会影响ID。如果给予的stack_id和当前资源栈的ID不一致，则返回400
	StackId *string `json:"stack_id,omitempty"`
}

func (o ContinueRollbackStackRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContinueRollbackStackRequestBody struct{}"
	}

	return strings.Join([]string{"ContinueRollbackStackRequestBody", string(data)}, " ")
}
