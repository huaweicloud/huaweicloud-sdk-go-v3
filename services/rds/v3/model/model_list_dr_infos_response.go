package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDrInfosResponse Response Object
type ListDrInfosResponse struct {

	// 总记录数
	TotalCount *int32 `json:"total_count,omitempty"`

	// 实例容灾信息
	InstanceDrInfos *[]InstanceDrInfo `json:"instance_dr_infos,omitempty"`
	HttpStatusCode  int               `json:"-"`
}

func (o ListDrInfosResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDrInfosResponse struct{}"
	}

	return strings.Join([]string{"ListDrInfosResponse", string(data)}, " ")
}
