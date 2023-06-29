package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataclassInfoRef Response of Dataclass Info Ref
type DataclassInfoRef struct {

	// Id value
	Id *string `json:"id,omitempty"`

	// Id value
	Name *string `json:"name,omitempty"`
}

func (o DataclassInfoRef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataclassInfoRef struct{}"
	}

	return strings.Join([]string{"DataclassInfoRef", string(data)}, " ")
}
