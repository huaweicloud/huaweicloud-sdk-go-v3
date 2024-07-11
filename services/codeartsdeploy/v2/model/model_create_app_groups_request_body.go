package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAppGroupsRequestBody 创建分组请求体
type CreateAppGroupsRequestBody struct {

	// 分组名称
	Name string `json:"name"`

	// 父分组id
	ParentId *string `json:"parent_id,omitempty"`
}

func (o CreateAppGroupsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppGroupsRequestBody struct{}"
	}

	return strings.Join([]string{"CreateAppGroupsRequestBody", string(data)}, " ")
}
