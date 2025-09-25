package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImageBasicImageResponse Response Object
type ListImageBasicImageResponse struct {

	// 总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 镜像的基础镜像信息
	DataList       *[]ImageBasicImageInfo `json:"data_list,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListImageBasicImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageBasicImageResponse struct{}"
	}

	return strings.Join([]string{"ListImageBasicImageResponse", string(data)}, " ")
}
