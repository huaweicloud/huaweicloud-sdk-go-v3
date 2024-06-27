package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateUserDefinedUrlKeyWordRequest Request Object
type CreateUserDefinedUrlKeyWordRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	Body *CreateBasicAwReq `json:"body,omitempty"`
}

func (o CreateUserDefinedUrlKeyWordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUserDefinedUrlKeyWordRequest struct{}"
	}

	return strings.Join([]string{"CreateUserDefinedUrlKeyWordRequest", string(data)}, " ")
}
