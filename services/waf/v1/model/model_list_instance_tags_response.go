package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceTagsResponse Response Object
type ListInstanceTagsResponse struct {

	// 独享引擎实例的标签信息
	Tags           *[]InstanceTag `json:"tags,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListInstanceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceTagsResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceTagsResponse", string(data)}, " ")
}
