package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAttachmentsResponse Response Object
type ListAttachmentsResponse struct {

	// 实际的数据类型：单个对象，集合 或 NULL
	Value          *[]AttachmentVo `json:"value,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListAttachmentsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAttachmentsResponse struct{}"
	}

	return strings.Join([]string{"ListAttachmentsResponse", string(data)}, " ")
}
