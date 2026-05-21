package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeletePoolsResponse Response Object
type BatchDeletePoolsResponse struct {

	// **参数解释**：请求ID。  **取值范围**：由数字、小写字母和中划线（-）组成的字符串，自动生成。
	RequestId *string `json:"request_id,omitempty"`

	// 后端服务器组批量删除后的响应结果。
	Pools          *[]BatchDeletePoolsResp `json:"pools,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o BatchDeletePoolsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeletePoolsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeletePoolsResponse", string(data)}, " ")
}
