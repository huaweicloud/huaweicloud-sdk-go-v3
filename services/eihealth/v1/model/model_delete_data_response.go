package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteDataResponse struct {

	// 数据作业ID
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDataResponse struct{}"
	}

	return strings.Join([]string{"DeleteDataResponse", string(data)}, " ")
}
