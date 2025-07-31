package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListContainerImagesResponse Response Object
type ListContainerImagesResponse struct {

	// 总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 容器镜像信息列表
	DataList       *[]ContainerImageInfo `json:"data_list,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListContainerImagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListContainerImagesResponse struct{}"
	}

	return strings.Join([]string{"ListContainerImagesResponse", string(data)}, " ")
}
