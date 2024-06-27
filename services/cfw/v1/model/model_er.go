package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Er struct {

	// ER ID
	ErId *string `json:"er_id,omitempty"`

	// ER连接ID
	ErAttachId *string `json:"er_attach_id,omitempty"`
}

func (o Er) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Er struct{}"
	}

	return strings.Join([]string{"Er", string(data)}, " ")
}
