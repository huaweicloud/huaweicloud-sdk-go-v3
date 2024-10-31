package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteServiceItemResponseBodyData struct {

	// 服务组成员id
	Id *string `json:"id,omitempty"`

	// 服务组成员名称，为源和目的端口
	Name *string `json:"name,omitempty"`
}

func (o DeleteServiceItemResponseBodyData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServiceItemResponseBodyData struct{}"
	}

	return strings.Join([]string{"DeleteServiceItemResponseBodyData", string(data)}, " ")
}
