package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteResourceViewsRequest Request Object
type DeleteResourceViewsRequest struct {

	// 视图ID。长度1~32之间。
	Id string `json:"id"`
}

func (o DeleteResourceViewsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResourceViewsRequest struct{}"
	}

	return strings.Join([]string{"DeleteResourceViewsRequest", string(data)}, " ")
}
