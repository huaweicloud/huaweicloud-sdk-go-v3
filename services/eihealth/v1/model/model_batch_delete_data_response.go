package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteDataResponse Response Object
type BatchDeleteDataResponse struct {

	// 数据作业ID
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchDeleteDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteDataResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteDataResponse", string(data)}, " ")
}
