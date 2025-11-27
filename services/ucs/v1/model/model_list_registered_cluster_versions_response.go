package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRegisteredClusterVersionsResponse Response Object
type ListRegisteredClusterVersionsResponse struct {

	// 集群版本列表
	Versions       *[]string `json:"versions,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListRegisteredClusterVersionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRegisteredClusterVersionsResponse struct{}"
	}

	return strings.Join([]string{"ListRegisteredClusterVersionsResponse", string(data)}, " ")
}
