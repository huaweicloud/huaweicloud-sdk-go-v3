package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RunSearchPictureResponse struct {

	// 搜索结果总数。
	Count *int32 `json:"count,omitempty"`

	// 搜索结果详情。
	Result         *[]SearchPictureItem `json:"result,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o RunSearchPictureResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunSearchPictureResponse struct{}"
	}

	return strings.Join([]string{"RunSearchPictureResponse", string(data)}, " ")
}
