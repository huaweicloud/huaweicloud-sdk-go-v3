package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListNextflowVersionResponse struct {

	// 版本总数
	Count *int32 `json:"count,omitempty"`

	// 版本列表
	Versions       *[]string `json:"versions,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListNextflowVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNextflowVersionResponse struct{}"
	}

	return strings.Join([]string{"ListNextflowVersionResponse", string(data)}, " ")
}
