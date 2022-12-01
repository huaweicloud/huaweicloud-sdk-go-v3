package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Query Service Set Detail Response Body
type ServiceSetDetailResponseDto struct {

	// 服务组id
	Id *string `json:"id,omitempty"`

	// 服务组名称
	Name string `json:"name"`

	// 服务组描述信息
	Description *string `json:"description,omitempty"`
}

func (o ServiceSetDetailResponseDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceSetDetailResponseDto struct{}"
	}

	return strings.Join([]string{"ServiceSetDetailResponseDto", string(data)}, " ")
}
