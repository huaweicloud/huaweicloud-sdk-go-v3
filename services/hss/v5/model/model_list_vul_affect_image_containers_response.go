package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVulAffectImageContainersResponse Response Object
type ListVulAffectImageContainersResponse struct {

	// 总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 列表
	DataList       *[]VulAffectImageContainersResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                                     `json:"-"`
}

func (o ListVulAffectImageContainersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVulAffectImageContainersResponse struct{}"
	}

	return strings.Join([]string{"ListVulAffectImageContainersResponse", string(data)}, " ")
}
