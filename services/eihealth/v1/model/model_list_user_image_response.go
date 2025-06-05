package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserImageResponse Response Object
type ListUserImageResponse struct {

	// 镜像总数
	Count *int32 `json:"count,omitempty"`

	// 镜像详情列表
	Images         *[]ImageDetailRsp `json:"images,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListUserImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserImageResponse struct{}"
	}

	return strings.Join([]string{"ListUserImageResponse", string(data)}, " ")
}
