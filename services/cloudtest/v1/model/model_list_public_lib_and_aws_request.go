package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPublicLibAndAwsRequest Request Object
type ListPublicLibAndAwsRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`
}

func (o ListPublicLibAndAwsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPublicLibAndAwsRequest struct{}"
	}

	return strings.Join([]string{"ListPublicLibAndAwsRequest", string(data)}, " ")
}
