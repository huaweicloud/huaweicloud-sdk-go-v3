package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVersionDetailsResponse Response Object
type ListVersionDetailsResponse struct {

	// VPC终端节点版本信息列表。
	Versions       *[]VersionObject `json:"versions,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListVersionDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVersionDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListVersionDetailsResponse", string(data)}, " ")
}
