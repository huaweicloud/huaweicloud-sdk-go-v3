package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteMetadataResponse struct {

	// 系统提示信息，执行成功时，字段可能为空。执行失败时，用于显示错误信息。
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage"`

	// 系统提示信息，执行成功时，字段可能为空。执行失败时，用于显示错误码。
	ErrorCode      *string `json:"errorCode,omitempty" xml:"errorCode"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteMetadataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMetadataResponse struct{}"
	}

	return strings.Join([]string{"DeleteMetadataResponse", string(data)}, " ")
}
