package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ServiceSetId struct {

	// 服务组Id
	Id *string `json:"id,omitempty"`

	// 服务组名称
	Name *string `json:"name,omitempty"`
}

func (o ServiceSetId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceSetId struct{}"
	}

	return strings.Join([]string{"ServiceSetId", string(data)}, " ")
}
