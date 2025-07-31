package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListK8sPodsResponse Response Object
type ListK8sPodsResponse struct {

	// pod总数
	TotalNum *int64 `json:"total_num,omitempty"`

	// pod列表
	DataList *[]PodBaseInfo `json:"data_list,omitempty"`

	// 实例基本信息列表
	PodInfoList    *[]ServerlessPodInfo `json:"pod_info_list,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListK8sPodsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListK8sPodsResponse struct{}"
	}

	return strings.Join([]string{"ListK8sPodsResponse", string(data)}, " ")
}
