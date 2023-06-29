package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAlertDetailDataclassRef dataclass对象
type ShowAlertDetailDataclassRef struct {

	// Id value
	Id *string `json:"id,omitempty"`

	// The name, display only
	Name *string `json:"name,omitempty"`
}

func (o ShowAlertDetailDataclassRef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlertDetailDataclassRef struct{}"
	}

	return strings.Join([]string{"ShowAlertDetailDataclassRef", string(data)}, " ")
}
