package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddImageMemberRequestBody 共享镜像成员请求体
type AddImageMemberRequestBody struct {

	// 被共享账号的PROJECT_ID
	Member *string `json:"member,omitempty"`
}

func (o AddImageMemberRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddImageMemberRequestBody struct{}"
	}

	return strings.Join([]string{"AddImageMemberRequestBody", string(data)}, " ")
}
