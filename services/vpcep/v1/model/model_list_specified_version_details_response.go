package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListSpecifiedVersionDetailsResponse struct {
	// VPC终端节点版本信息列表。

	Versions       *[]Versions `json:"versions,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListSpecifiedVersionDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSpecifiedVersionDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListSpecifiedVersionDetailsResponse", string(data)}, " ")
}
