package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResultValueStringForOk struct {

	// ok
	Value *string `json:"value,omitempty"`
}

func (o ResultValueStringForOk) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResultValueStringForOk struct{}"
	}

	return strings.Join([]string{"ResultValueStringForOk", string(data)}, " ")
}
