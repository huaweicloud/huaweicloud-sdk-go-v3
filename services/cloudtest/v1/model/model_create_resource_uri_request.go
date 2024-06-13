package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateResourceUriRequest Request Object
type CreateResourceUriRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`
}

func (o CreateResourceUriRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourceUriRequest struct{}"
	}

	return strings.Join([]string{"CreateResourceUriRequest", string(data)}, " ")
}
