package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAllInstanceRepositoriesResponse Response Object
type ListAllInstanceRepositoriesResponse struct {
	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 仓库列表详情
	Repositories   *[]InstanceRepository `json:"repositories,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListAllInstanceRepositoriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllInstanceRepositoriesResponse struct{}"
	}

	return strings.Join([]string{"ListAllInstanceRepositoriesResponse", string(data)}, " ")
}
