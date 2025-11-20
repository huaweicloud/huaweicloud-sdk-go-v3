package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MigrateDnInfoOpenResponse struct {

	// rds实例id。
	DnInstanceId *string `json:"dn_instance_id,omitempty"`

	// rds实例名称。
	DnInstanceName *string `json:"dn_instance_name,omitempty"`
}

func (o MigrateDnInfoOpenResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateDnInfoOpenResponse struct{}"
	}

	return strings.Join([]string{"MigrateDnInfoOpenResponse", string(data)}, " ")
}
