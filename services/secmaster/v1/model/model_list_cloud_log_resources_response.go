package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCloudLogResourcesResponse Response Object
type ListCloudLogResourcesResponse struct {

	// 数据集列表
	Datasets *[]DatasetItem `json:"datasets,omitempty"`

	// 表示资源是否存在
	Exist *bool `json:"exist,omitempty"`

	// 工作空间列表
	Workspaces     *[]string `json:"workspaces,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListCloudLogResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudLogResourcesResponse struct{}"
	}

	return strings.Join([]string{"ListCloudLogResourcesResponse", string(data)}, " ")
}
