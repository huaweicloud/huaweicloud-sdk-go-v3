package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAllPopsResponse Response Object
type ListAllPopsResponse struct {

	// 接入点列表。
	Pops *[]PopOuterDetail `json:"pops,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListAllPopsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllPopsResponse struct{}"
	}

	return strings.Join([]string{"ListAllPopsResponse", string(data)}, " ")
}
