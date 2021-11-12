package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type TagSatisfactionResponse struct {
	// 调用成功时的返回请求ID。  调用失败时无此字段。

	RequestId *string `json:"request_id,omitempty"`
	// 反馈满意度的时间。格式为“yyyy-MM-dd THH:mm:ssZ”。其中，T指某个时间的开始；Z指UTC时间。  调用失败时无此字段。

	UpdatedTime    *string `json:"updated_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o TagSatisfactionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagSatisfactionResponse struct{}"
	}

	return strings.Join([]string{"TagSatisfactionResponse", string(data)}, " ")
}
