package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeletePoolsResp 批量删除的响应结果
type BatchDeletePoolsResp struct {

	// 被删除的后端服务器组id。
	Id string `json:"id"`

	// 对应id的后端服务器组删除后的结果，not found表示后端服务器组不存在，successful表示删除成功
	RetStatus string `json:"ret_status"`

	// 错误码，删除失败时返回此字段
	RetCode *string `json:"ret_code,omitempty"`
}

func (o BatchDeletePoolsResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeletePoolsResp struct{}"
	}

	return strings.Join([]string{"BatchDeletePoolsResp", string(data)}, " ")
}
