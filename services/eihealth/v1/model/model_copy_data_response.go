package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CopyDataResponse Response Object
type CopyDataResponse struct {

	// 数据作业ID
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CopyDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyDataResponse struct{}"
	}

	return strings.Join([]string{"CopyDataResponse", string(data)}, " ")
}
