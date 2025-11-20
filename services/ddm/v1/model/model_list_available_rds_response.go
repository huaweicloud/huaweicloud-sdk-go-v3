package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAvailableRdsResponse Response Object
type ListAvailableRdsResponse struct {

	// 可用后端DN信息。
	DataNodes      *[]AvailableDnInstance `json:"data_nodes,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListAvailableRdsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailableRdsResponse struct{}"
	}

	return strings.Join([]string{"ListAvailableRdsResponse", string(data)}, " ")
}
