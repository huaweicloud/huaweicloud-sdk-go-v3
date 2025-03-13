package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportSecurityBuiltinCategoryGroupsResponse Response Object
type ImportSecurityBuiltinCategoryGroupsResponse struct {

	// 导入结果。
	Result         *[]ImportDefaultCategoryResultDto `json:"result,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o ImportSecurityBuiltinCategoryGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportSecurityBuiltinCategoryGroupsResponse struct{}"
	}

	return strings.Join([]string{"ImportSecurityBuiltinCategoryGroupsResponse", string(data)}, " ")
}
