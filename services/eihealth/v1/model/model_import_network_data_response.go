package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportNetworkDataResponse Response Object
type ImportNetworkDataResponse struct {

	// 数据作业ID
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ImportNetworkDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportNetworkDataResponse struct{}"
	}

	return strings.Join([]string{"ImportNetworkDataResponse", string(data)}, " ")
}
