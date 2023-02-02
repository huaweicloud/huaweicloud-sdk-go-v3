package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListStackEventsRequest struct {

	// 用户指定的，对于此请求的唯一ID，用于定位某个请求，推荐使用UUID
	ClientRequestId string `json:"Client-Request-Id"`

	// 用户希望操作的资源栈名称
	StackName string `json:"stack_name"`

	// 用户希望描述的资源栈的Id。若stack_name和stack_id同时存在，则资源编排服务会检查是否两个匹配，否则返回400
	StackId *string `json:"stack_id,omitempty"`

	// 部署时API返回的id（uuid）
	DeploymentId *string `json:"deployment_id,omitempty"`

	// 过滤条件  * 与（AND）运算符使用逗号（，）定义 * 或（OR）运算符使用竖线（|）定义，OR运算符优先级高于AND运算符 * 不支持括号 * 过滤运算符仅支持等号（==） * 过滤参数名及其值仅支持包含大小写英文、数字和下划线 * 过滤条件中禁止使用分号，若有分号，则此条过滤会被忽略 * 一个过滤参数仅能与一个与条件相关，一个与条件中的多个或条件仅能与一个过滤参数相关
	Filter *string `json:"filter,omitempty"`

	// 选择的属性名称  * 属性名仅支持包含大小写英文、数字和下划线 * 多个属性名称之间以逗号（，）分隔
	Field *string `json:"field,omitempty"`
}

func (o ListStackEventsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStackEventsRequest struct{}"
	}

	return strings.Join([]string{"ListStackEventsRequest", string(data)}, " ")
}
