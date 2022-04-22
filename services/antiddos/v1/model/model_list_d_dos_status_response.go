package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListDDosStatusResponse struct {

	// 弹性IP总数
	Total *int32 `json:"total,omitempty"`

	// 防护状态列表
	DdosStatus     *[]DDosStatus `json:"ddosStatus,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListDDosStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDDosStatusResponse struct{}"
	}

	return strings.Join([]string{"ListDDosStatusResponse", string(data)}, " ")
}
