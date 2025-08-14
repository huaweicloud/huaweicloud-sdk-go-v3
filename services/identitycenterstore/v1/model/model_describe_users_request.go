package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DescribeUsersRequest Request Object
type DescribeUsersRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// 身份源的全局唯一标识符（ID）
	IdentityStoreId string `json:"identity_store_id"`

	Body *DescribeUsersRequestBody `json:"body,omitempty"`
}

func (o DescribeUsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DescribeUsersRequest struct{}"
	}

	return strings.Join([]string{"DescribeUsersRequest", string(data)}, " ")
}
