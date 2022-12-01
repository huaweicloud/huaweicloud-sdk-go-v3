package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RunQueryCustomTagsResponse struct {

	// 查询/删除自定义标签结果
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunQueryCustomTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunQueryCustomTagsResponse struct{}"
	}

	return strings.Join([]string{"RunQueryCustomTagsResponse", string(data)}, " ")
}
