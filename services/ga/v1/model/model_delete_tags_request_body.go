package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// delete tag request
type DeleteTagsRequestBody struct {

	// 标签列表。
	Tags []DeletingResourceTag `json:"tags"`
}

func (o DeleteTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTagsRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteTagsRequestBody", string(data)}, " ")
}
