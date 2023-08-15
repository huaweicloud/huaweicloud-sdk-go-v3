package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLiveDataApiDeploymentHistoryV2Response Response Object
type ListLiveDataApiDeploymentHistoryV2Response struct {

	// 本次返回的列表长度
	Size int32 `json:"size"`

	// 满足条件的记录数
	Total int64 `json:"total"`

	// 本次查询到的后端API部署结果对象列表
	Histories      *[]LdApiDeployHistoryInfo `json:"histories,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListLiveDataApiDeploymentHistoryV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLiveDataApiDeploymentHistoryV2Response struct{}"
	}

	return strings.Join([]string{"ListLiveDataApiDeploymentHistoryV2Response", string(data)}, " ")
}
