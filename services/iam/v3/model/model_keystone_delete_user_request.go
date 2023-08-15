package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneDeleteUserRequest Request Object
type KeystoneDeleteUserRequest struct {

	// 待删除的IAM用户ID，获取方式请参见：[获取用户ID](https://support.huaweicloud.com/api-iam/iam_17_0002.html)。
	UserId string `json:"user_id"`
}

func (o KeystoneDeleteUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneDeleteUserRequest struct{}"
	}

	return strings.Join([]string{"KeystoneDeleteUserRequest", string(data)}, " ")
}
