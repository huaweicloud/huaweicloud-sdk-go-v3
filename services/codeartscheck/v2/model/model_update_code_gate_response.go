package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCodeGateResponse Response Object
type UpdateCodeGateResponse struct {

	// **参数解释：** 更新成功时返回\"更新门禁配置成功\"。 **取值范围：** 不涉及。
	Result *string `json:"result,omitempty"`

	// **参数解释：** HTTP响应状态。 **取值范围：** - OK：请求已成功处理，服务器返回了请求的资源或操作结果。 - ACCEPTED：请求已被服务器接受，但处理尚未完成（通常用于异步操作）。 - INTERNAL_SERVER_ERROR：接口内部异常。 - BAD_REQUEST：参数校验异常。
	HttpStatus *string `json:"httpStatus,omitempty"`

	// **参数解释：** 接口响应状态。 **取值范围：** - success：接口响应成功。 - error：接口响应失败。
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateCodeGateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCodeGateResponse struct{}"
	}

	return strings.Join([]string{"UpdateCodeGateResponse", string(data)}, " ")
}
