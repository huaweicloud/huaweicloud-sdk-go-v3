package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSeveritiesResponse Response Object
type ListSeveritiesResponse struct {

	// 是否展示
	Show *bool `json:"show,omitempty"`

	// 严重性列表
	SeverityList   *[]SeverityV2Do `json:"severity_list,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListSeveritiesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSeveritiesResponse struct{}"
	}

	return strings.Join([]string{"ListSeveritiesResponse", string(data)}, " ")
}
