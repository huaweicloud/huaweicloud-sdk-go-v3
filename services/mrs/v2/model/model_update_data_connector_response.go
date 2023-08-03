package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDataConnectorResponse Response Object
type UpdateDataConnectorResponse struct {

	// 更新映射请求操作结果，succeeded为操作成功，failed为操作失败
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateDataConnectorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDataConnectorResponse struct{}"
	}

	return strings.Join([]string{"UpdateDataConnectorResponse", string(data)}, " ")
}
