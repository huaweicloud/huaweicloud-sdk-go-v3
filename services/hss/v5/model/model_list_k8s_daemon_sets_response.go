package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListK8sDaemonSetsResponse Response Object
type ListK8sDaemonSetsResponse struct {

	// daemonset总量
	TotalNum *int32 `json:"total_num,omitempty"`

	// daemonset基本信息列表
	DataList       *[]DaemonSetInfo `json:"data_list,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListK8sDaemonSetsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListK8sDaemonSetsResponse struct{}"
	}

	return strings.Join([]string{"ListK8sDaemonSetsResponse", string(data)}, " ")
}
