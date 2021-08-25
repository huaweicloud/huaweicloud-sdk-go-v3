package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type TagLaborResponse struct {
	// 调用成功时的返回请求ID。  调用失败时无此字段。

	RequestId *string `json:"request_id,omitempty"`
	// 转人工的时间。格式为“yyyy-MM-dd THH:mm:ssZ”。其中，T指某个时间的开始；Z指UTC时间。  调用失败时无此字段。

	UpdatedTime    *string `json:"updated_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o TagLaborResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "TagLaborResponse struct{}"
	}

	return strings.Join([]string{"TagLaborResponse", string(data)}, " ")
}
