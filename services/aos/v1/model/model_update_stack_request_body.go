package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateStackRequestBody struct {

	// 资源栈的描述。可用于客户识别自己的资源栈。
	Description *string `json:"description,omitempty"`

	// 资源栈（stack）的唯一Id。  此Id由资源编排服务在生成资源栈的时候生成，为UUID。  由于资源栈名仅仅在同一时间下唯一，即用户允许先生成一个叫HelloWorld的资源栈，删除，再重新创建一个同名资源栈。  对于团队并行开发，用户可能希望确保，当前我操作的资源栈就是我认为的那个，而不是其他队友删除后创建的同名资源栈。因此，使用ID就可以做到强匹配。  资源编排服务保证每次创建的资源栈所对应的ID都不相同，更新不会影响ID。如果给与的stack_id和当前资源栈的ID不一致，则返回400
	StackId *string `json:"stack_id,omitempty"`

	// 删除保护的标识位，如果不传默认为false，即默认不开启资源栈删除保护（删除保护开启后资源栈不允许被删除）  *在UpdateStack API中，若该参数未在RequestBody中给予，则不会对资源栈的删除保护属性进行更新*
	EnableDeletionProtection *bool `json:"enable_deletion_protection,omitempty"`

	// 自动回滚的标识位，如果不传默认为false，即默认不开启资源栈自动回滚（自动回滚开启后，如果部署失败，则会自动回滚，并返回上一个稳定状态）  *在UpdateStack API中，若该参数未在RequestBody中给予，则不会对资源栈的自动回滚属性进行更新*
	EnableAutoRollback *bool `json:"enable_auto_rollback,omitempty"`

	// 委托授权的信息。
	Agencies *[]Agency `json:"agencies,omitempty"`
}

func (o UpdateStackRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStackRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateStackRequestBody", string(data)}, " ")
}
