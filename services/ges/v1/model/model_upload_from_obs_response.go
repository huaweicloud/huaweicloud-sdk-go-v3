package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadFromObsResponse Response Object
type UploadFromObsResponse struct {

	// 系统提示信息，执行成功时，字段可能为空。执行失败时，用于显示错误信息。
	ErrorMessage *string `json:"errorMessage,omitempty"`

	// 系统提示信息，执行成功时，字段可能为空。执行失败时，用于显示错误码。
	ErrorCode *string `json:"errorCode,omitempty"`

	// 元数据的id。
	Id *string `json:"id,omitempty"`

	// 元数据的名字。
	Name           *string `json:"name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UploadFromObsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadFromObsResponse struct{}"
	}

	return strings.Join([]string{"UploadFromObsResponse", string(data)}, " ")
}
