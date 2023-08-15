package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInstanceByIdRequest Request Object
type DeleteInstanceByIdRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`
}

func (o DeleteInstanceByIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstanceByIdRequest struct{}"
	}

	return strings.Join([]string{"DeleteInstanceByIdRequest", string(data)}, " ")
}
