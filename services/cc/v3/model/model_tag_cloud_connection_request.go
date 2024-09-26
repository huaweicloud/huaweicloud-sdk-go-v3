package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagCloudConnectionRequest Request Object
type TagCloudConnectionRequest struct {

	// 实例ID。
	Id string `json:"id"`

	Body *TagCloudConnectionRequestBody `json:"body,omitempty"`
}

func (o TagCloudConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagCloudConnectionRequest struct{}"
	}

	return strings.Join([]string{"TagCloudConnectionRequest", string(data)}, " ")
}
