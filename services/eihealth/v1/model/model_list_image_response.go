package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImageResponse Response Object
type ListImageResponse struct {

	// 镜像总数
	Count *int32 `json:"count,omitempty"`

	// 镜像详情列表
	Images         *[]ImageDetailRsp `json:"images,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageResponse struct{}"
	}

	return strings.Join([]string{"ListImageResponse", string(data)}, " ")
}
