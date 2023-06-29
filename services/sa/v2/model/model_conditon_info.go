package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConditonInfo ConditonInfo
type ConditonInfo struct {

	// name
	Name *string `json:"name,omitempty"`

	// search data
	Data *[]string `json:"data,omitempty"`
}

func (o ConditonInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConditonInfo struct{}"
	}

	return strings.Join([]string{"ConditonInfo", string(data)}, " ")
}
