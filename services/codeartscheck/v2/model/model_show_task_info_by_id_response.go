package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaskInfoByIdResponse Response Object
type ShowTaskInfoByIdResponse struct {

	// **参数解释：** HTTP响应状态。 **取值范围：** - OK：请求已成功处理，服务器返回了请求的资源或操作结果。 - ACCEPTED：请求已被服务器接受，但处理尚未完成（通常用于异步操作）。 - INTERNAL_SERVER_ERROR：接口内部异常。 - BAD_REQUEST：参数校验异常。
	HttpStatus *string `json:"httpStatus,omitempty"`

	Result *TaskBasicInfoV3 `json:"result,omitempty"`

	// **参数解释：** 接口响应状态。 **取值范围：** - success：接口响应成功。 - error：接口响应失败。
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowTaskInfoByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskInfoByIdResponse struct{}"
	}

	return strings.Join([]string{"ShowTaskInfoByIdResponse", string(data)}, " ")
}
