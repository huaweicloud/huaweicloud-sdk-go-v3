package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListImageByTagsResponse struct {

	// 镜像信息列表
	Resources *[]ShowImageByTagsResource `json:"resources,omitempty" xml:"resources"`

	// 总记录数
	TotalCount     *int32 `json:"total_count,omitempty" xml:"total_count"`
	HttpStatusCode int    `json:"-"`
}

func (o ListImageByTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageByTagsResponse struct{}"
	}

	return strings.Join([]string{"ListImageByTagsResponse", string(data)}, " ")
}
