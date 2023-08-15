package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSharedTagsRequest Request Object
type ShowSharedTagsRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type"`

	// 共享ID
	ShareId string `json:"share_id"`
}

func (o ShowSharedTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSharedTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowSharedTagsRequest", string(data)}, " ")
}
