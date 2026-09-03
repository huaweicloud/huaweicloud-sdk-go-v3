package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IpGroup struct {

	// IP组名称
	GroupName *string `json:"group_name,omitempty"`

	// IP组id
	Id *string `json:"id,omitempty"`
}

func (o IpGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpGroup struct{}"
	}

	return strings.Join([]string{"IpGroup", string(data)}, " ")
}
