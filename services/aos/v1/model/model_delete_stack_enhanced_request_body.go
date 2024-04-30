package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteStackEnhancedRequestBody struct {

	// 资源栈（stack）的唯一ID。  此ID由资源编排服务在生成资源栈的时候生成，为UUID。  由于资源栈名仅仅在同一时间下唯一，即用户允许先生成一个叫HelloWorld的资源栈，删除，再重新创建一个同名资源栈。  对于团队并行开发，用户可能希望确保，当前我操作的资源栈就是我认为的那个，而不是其他队友删除后创建的同名资源栈。因此，使用ID就可以做到强匹配。  资源编排服务保证每次创建的资源栈所对应的ID都不相同，更新不会影响ID。如果给予的stack_id和当前资源栈的ID不一致，则返回400
	StackId *string `json:"stack_id,omitempty"`

	// 删除资源栈是否保留资源的标志位，如果不传默认为false，即默认不保留资源（删除资源栈后会删除资源栈中的资源）  * DeleteStackEnhanced API中，如果该参数未在RequestBody中给予，则删除时不会保留资源栈中的资源*
	RetainAllResources *bool `json:"retain_all_resources,omitempty"`
}

func (o DeleteStackEnhancedRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteStackEnhancedRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteStackEnhancedRequestBody", string(data)}, " ")
}
