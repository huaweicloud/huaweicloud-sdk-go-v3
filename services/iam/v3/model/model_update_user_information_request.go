package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateUserInformationRequest struct {
	// 待修改信息的IAM用户ID，获取方式请参见：[获取账号、IAM用户、项目、用户组、委托的名称和ID](https://support.huaweicloud.com/api-iam/iam_17_0002.html)。

	UserId string `json:"user_id"`

	Body *UpdateUserInformationRequestBody `json:"body,omitempty"`
}

func (o UpdateUserInformationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserInformationRequest struct{}"
	}

	return strings.Join([]string{"UpdateUserInformationRequest", string(data)}, " ")
}
