package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowProjectDetailsAndStatusRequest struct {
	// 待查询的项目ID，获取方式请参见：[获取账号、IAM用户、项目、用户组、委托的名称和ID](https://support.huaweicloud.com/api-iam/iam_17_0002.html)。

	ProjectId string `json:"project_id"`
}

func (o ShowProjectDetailsAndStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectDetailsAndStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowProjectDetailsAndStatusRequest", string(data)}, " ")
}
