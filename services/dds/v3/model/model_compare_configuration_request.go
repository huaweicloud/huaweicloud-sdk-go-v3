package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CompareConfigurationRequest Request Object
type CompareConfigurationRequest struct {
	Body *DiffConfigurationRequest `json:"body,omitempty"`
}

func (o CompareConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareConfigurationRequest struct{}"
	}

	return strings.Join([]string{"CompareConfigurationRequest", string(data)}, " ")
}
