package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResultValueString Data object returned by the request.
type ResultValueString struct {
	Value *string `json:"value,omitempty"`
}

func (o ResultValueString) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResultValueString struct{}"
	}

	return strings.Join([]string{"ResultValueString", string(data)}, " ")
}
