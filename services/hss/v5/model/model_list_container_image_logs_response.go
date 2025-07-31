package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListContainerImageLogsResponse Response Object
type ListContainerImageLogsResponse struct {

	// 日志总条数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 容器镜像使用日志列表
	DataList       *[]ContainerImageLogResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o ListContainerImageLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListContainerImageLogsResponse struct{}"
	}

	return strings.Join([]string{"ListContainerImageLogsResponse", string(data)}, " ")
}
