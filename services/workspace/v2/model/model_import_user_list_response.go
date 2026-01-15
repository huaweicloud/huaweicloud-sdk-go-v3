package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportUserListResponse Response Object
type ImportUserListResponse struct {

	// 错误码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述。
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 错误详情。
	ErrorDetail *string `json:"error_detail,omitempty"`

	// 加密后的详细拒绝原因，用户可以自行调用STS服务的decode-authorization-message接口进行解密。
	EncodedAuthorizationMessage *string `json:"encoded_authorization_message,omitempty"`

	// 成功列表。
	UserDetailList *[]DesktopUserDetail `json:"user_detail_list,omitempty"`

	// 失败列表。
	FailedDetailList *[]DesktopUserDetail `json:"failed_detail_list,omitempty"`

	// 用户列表数量。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ImportUserListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportUserListResponse struct{}"
	}

	return strings.Join([]string{"ImportUserListResponse", string(data)}, " ")
}
