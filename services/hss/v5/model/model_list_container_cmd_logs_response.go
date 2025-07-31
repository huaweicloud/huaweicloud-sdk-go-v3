package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListContainerCmdLogsResponse Response Object
type ListContainerCmdLogsResponse struct {

	// 日志总条数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 容器内运行命令的日志列表
	DataList       *[]ContainerCmdLogResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ListContainerCmdLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListContainerCmdLogsResponse struct{}"
	}

	return strings.Join([]string{"ListContainerCmdLogsResponse", string(data)}, " ")
}
