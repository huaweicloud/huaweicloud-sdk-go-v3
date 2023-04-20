package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAppImageResponse struct {

	// 总记录数
	Count *int32 `json:"count,omitempty"`

	PageInfo *PageInfoDto `json:"page_info,omitempty"`

	// 每页记录数
	AppImages      *[]QueryAppImageResponseDto `json:"app_images,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListAppImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppImageResponse struct{}"
	}

	return strings.Join([]string{"ListAppImageResponse", string(data)}, " ")
}
