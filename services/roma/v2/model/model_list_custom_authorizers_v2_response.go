package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListCustomAuthorizersV2Response struct {
	// 本次返回的列表长度

	Size int32 `json:"size"`
	// 满足条件的记录数

	Total int64 `json:"total"`
	// 自定义认证列表

	AuthorizerList *[]AuthorizerResp `json:"authorizer_list,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListCustomAuthorizersV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomAuthorizersV2Response struct{}"
	}

	return strings.Join([]string{"ListCustomAuthorizersV2Response", string(data)}, " ")
}
