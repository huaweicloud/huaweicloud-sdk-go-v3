package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeployInstanceResponse Response Object
type DeployInstanceResponse struct {

	// 实例id
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeployInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeployInstanceResponse struct{}"
	}

	return strings.Join([]string{"DeployInstanceResponse", string(data)}, " ")
}
