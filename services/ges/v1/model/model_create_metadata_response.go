package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateMetadataResponse struct {

	// 系统提示信息，执行成功时，字段可能为空。执行失败时，用于显示错误信息。
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage"`

	// 系统提示信息，执行成功时，字段可能为空。执行失败时，用于显示错误码。
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode"`

	// 元数据ID。
	Id *string `json:"id,omitempty" xml:"id"`

	// 元数据名字。
	Name           *string `json:"name,omitempty" xml:"name"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateMetadataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMetadataResponse struct{}"
	}

	return strings.Join([]string{"CreateMetadataResponse", string(data)}, " ")
}
