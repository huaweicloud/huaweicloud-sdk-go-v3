package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListConnectionNumberData struct {

	// 连接数名称
	Name *string `json:"name,omitempty"`

	// items
	List *[]ListConnectionNumberDataList `json:"list,omitempty"`
}

func (o ListConnectionNumberData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConnectionNumberData struct{}"
	}

	return strings.Join([]string{"ListConnectionNumberData", string(data)}, " ")
}
