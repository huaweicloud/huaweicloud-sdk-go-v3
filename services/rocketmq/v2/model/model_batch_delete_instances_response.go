package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteInstancesResponse Response Object
type BatchDeleteInstancesResponse struct {

	// 修改实例的结果。
	Results        *[]BatchDeleteInstanceRespResults `json:"results,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o BatchDeleteInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteInstancesResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteInstancesResponse", string(data)}, " ")
}
