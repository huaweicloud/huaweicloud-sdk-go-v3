package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StackIdPrimitiveTypeHolder struct {

	// 资源栈（stack）的唯一Id。  此Id由RF在生成资源栈的时候生成，为UUID。  由于堆栈名仅仅在同一时间下唯一，即用户允许先生成一个叫HelloWorld的堆栈，删除，再重新创建一个同名堆栈。  对于团队并行开发，用户可能希望确保，当前我操作的堆栈就是我认为的那个，而不是其他队友删除后创建的同名堆栈。因此，使用ID就可以做到强匹配。  RF保证每次创建的资源栈所对应的ID都不相同，更新不会影响ID。如果给与的stack_id和当前资源栈的ID不一致，则返回400
	StackId *string `json:"stack_id,omitempty"`
}

func (o StackIdPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StackIdPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"StackIdPrimitiveTypeHolder", string(data)}, " ")
}
