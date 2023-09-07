package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateConnectionInfoResponse Response Object
type UpdateConnectionInfoResponse struct {

	// 状态码
	ResCode *int32 `json:"res_code,omitempty"`

	// 成功信息
	ResLog *string `json:"res_log,omitempty"`

	// 成功信息
	ResMsg         *string `json:"res_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateConnectionInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConnectionInfoResponse struct{}"
	}

	return strings.Join([]string{"UpdateConnectionInfoResponse", string(data)}, " ")
}
