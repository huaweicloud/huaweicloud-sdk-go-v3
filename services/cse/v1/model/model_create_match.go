package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateMatch struct {
	Headers *CreateMatchHeaders `json:"headers,omitempty"`
}

func (o CreateMatch) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMatch struct{}"
	}

	return strings.Join([]string{"CreateMatch", string(data)}, " ")
}
