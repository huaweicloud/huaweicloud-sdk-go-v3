package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLocalImageContainersResponse Response Object
type ListLocalImageContainersResponse struct {

	// 总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 查询本地镜像的容器列表
	DataList       *[]LocalImageContainerInfo `json:"data_list,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListLocalImageContainersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLocalImageContainersResponse struct{}"
	}

	return strings.Join([]string{"ListLocalImageContainersResponse", string(data)}, " ")
}
