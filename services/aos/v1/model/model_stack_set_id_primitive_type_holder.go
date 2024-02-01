package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StackSetIdPrimitiveTypeHolder struct {

	// 资源栈集（stack_set）的唯一ID。  此ID由资源编排服务在生成资源栈集的时候生成，为UUID。  由于资源栈集名仅仅在同一时间下唯一，即用户允许先生成一个叫HelloWorld的资源栈集，删除，在重新创建一个同名资源栈集。  对于团队并行开发，用户可能希望确保，当前我操作的资源栈集就是我以为的那个，而不是又其他队友删除后创建的同名资源栈集。因此，使用ID就可以做到强匹配。  资源编排服务保证每次创建的资源栈集所对应的ID都不相同，更新不会影响ID。如果给予的stack_set_id和当前资源栈集的ID不一致，则返回400
	StackSetId *string `json:"stack_set_id,omitempty"`
}

func (o StackSetIdPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StackSetIdPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"StackSetIdPrimitiveTypeHolder", string(data)}, " ")
}
