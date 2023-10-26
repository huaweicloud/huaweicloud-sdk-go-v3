package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBackendTargetsRequest Request Object
type ListBackendTargetsRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type"`

	// 文件系统id
	ShareId string `json:"share_id"`

	// 查询列表返回元素个数
	Limit *int32 `json:"limit,omitempty"`

	// 查询列表偏移量
	Marker *string `json:"marker,omitempty"`
}

func (o ListBackendTargetsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackendTargetsRequest struct{}"
	}

	return strings.Join([]string{"ListBackendTargetsRequest", string(data)}, " ")
}
