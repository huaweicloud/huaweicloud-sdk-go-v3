package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListMetadatasResponse struct {

	// 系统提示信息，执行成功时，字段可能为空。执行失败时，用于显示错误信息。
	ErrorMessage *string `json:"errorMessage,omitempty"`

	// 系统提示信息，执行成功时，字段可能为空。执行失败时，用于显示错误码。
	ErrorCode *string `json:"errorCode,omitempty"`

	// 元数据返回个数。请求失败时，字段为空。
	SchemaCount *int32 `json:"schemaCount,omitempty"`

	// 当前projectId下的所有元数据列表。请求失败时，字段为空。
	SchemaList     *[]Metadata `json:"schemaList,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListMetadatasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMetadatasResponse struct{}"
	}

	return strings.Join([]string{"ListMetadatasResponse", string(data)}, " ")
}
