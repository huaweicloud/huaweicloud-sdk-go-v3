package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListVpcProtectsRequest struct {

	// 防护对象id
	ObjectId string `json:"object_id"`
}

func (o ListVpcProtectsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVpcProtectsRequest struct{}"
	}

	return strings.Join([]string{"ListVpcProtectsRequest", string(data)}, " ")
}
