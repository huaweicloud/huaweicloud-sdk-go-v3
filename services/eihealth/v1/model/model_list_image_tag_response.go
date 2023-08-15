package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImageTagResponse Response Object
type ListImageTagResponse struct {

	// 镜像版本总数
	Count *int32 `json:"count,omitempty"`

	// 镜像版本详情列表
	Tags           *[]GetTagDetailRsp `json:"tags,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListImageTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageTagResponse struct{}"
	}

	return strings.Join([]string{"ListImageTagResponse", string(data)}, " ")
}
