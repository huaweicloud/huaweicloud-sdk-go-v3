package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSessionByUserNameRequest Request Object
type ListSessionByUserNameRequest struct {

	// 用户名。
	UserName string `json:"user_name"`
}

func (o ListSessionByUserNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSessionByUserNameRequest struct{}"
	}

	return strings.Join([]string{"ListSessionByUserNameRequest", string(data)}, " ")
}
