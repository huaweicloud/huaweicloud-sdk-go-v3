package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImageMembersRequest Request Object
type ListImageMembersRequest struct {

	// 镜像id。
	ImageId string `json:"image_id"`
}

func (o ListImageMembersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageMembersRequest struct{}"
	}

	return strings.Join([]string{"ListImageMembersRequest", string(data)}, " ")
}
