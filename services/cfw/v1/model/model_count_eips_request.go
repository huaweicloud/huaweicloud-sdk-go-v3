package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CountEipsRequest struct {

	// 防护对象ID
	ObjectId string `json:"object_id"`
}

func (o CountEipsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountEipsRequest struct{}"
	}

	return strings.Join([]string{"CountEipsRequest", string(data)}, " ")
}
