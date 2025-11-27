package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteApplicationComponentRequest Request Object
type DeleteApplicationComponentRequest struct {

	// 分组id。
	Id string `json:"id"`
}

func (o DeleteApplicationComponentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteApplicationComponentRequest struct{}"
	}

	return strings.Join([]string{"DeleteApplicationComponentRequest", string(data)}, " ")
}
