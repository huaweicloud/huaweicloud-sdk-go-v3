package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterPortResponse Response Object
type ListClusterPortResponse struct {

	// 端口总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 端口列表
	Result         *[]ElbClusterPortReponseBody `json:"result,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListClusterPortResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterPortResponse struct{}"
	}

	return strings.Join([]string{"ListClusterPortResponse", string(data)}, " ")
}
