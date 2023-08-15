package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteGetImagesListResponse Response Object
type ExecuteGetImagesListResponse struct {

	// 配额
	Quota int32 `json:"quota"`

	// 总数
	Total int32 `json:"total"`

	// 偏移
	Offset int32 `json:"offset"`

	// 返回数量
	Count int32 `json:"count"`

	// 图片列表
	Images         []Image `json:"images"`
	HttpStatusCode int     `json:"-"`
}

func (o ExecuteGetImagesListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteGetImagesListResponse struct{}"
	}

	return strings.Join([]string{"ExecuteGetImagesListResponse", string(data)}, " ")
}
