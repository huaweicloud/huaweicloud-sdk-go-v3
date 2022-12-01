package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RunDeleteCustomTagsResponse struct {

	// 查询/删除自定义标签结果
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunDeleteCustomTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunDeleteCustomTagsResponse struct{}"
	}

	return strings.Join([]string{"RunDeleteCustomTagsResponse", string(data)}, " ")
}
