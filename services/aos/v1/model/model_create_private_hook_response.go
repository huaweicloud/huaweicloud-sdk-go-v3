package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePrivateHookResponse Response Object
type CreatePrivateHookResponse struct {

	// 私有hook（private-hook）的唯一Id。  此Id由资源编排服务在生成私有hook的时候生成，为UUID。  由于私有hook名称仅仅在同一时间下唯一，即用户允许先生成一个叫HelloWorld的私有hook，删除，再重新创建一个同名私有hook。  对于团队并行开发，用户可能希望确保，当前我操作的私有hook就是我认为的那个，而不是其他队友删除后创建的同名私有hook。因此，使用Id就可以做到强匹配。  资源编排服务保证每次创建的私有hook所对应的Id都不相同，更新不会影响Id。如果给与的hook_id和当前hook的Id不一致，则返回400。
	HookId         *string `json:"hook_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePrivateHookResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrivateHookResponse struct{}"
	}

	return strings.Join([]string{"CreatePrivateHookResponse", string(data)}, " ")
}
