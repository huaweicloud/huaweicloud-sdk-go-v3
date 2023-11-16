package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBridgeFunctionsResponse Response Object
type ListBridgeFunctionsResponse struct {

	// 函数绑定的servicebridge函数列表。
	Body           *[]ListFunctionResult `json:"body,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListBridgeFunctionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBridgeFunctionsResponse struct{}"
	}

	return strings.Join([]string{"ListBridgeFunctionsResponse", string(data)}, " ")
}
