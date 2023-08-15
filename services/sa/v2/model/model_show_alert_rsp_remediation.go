package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAlertRspRemediation 补救措施
type ShowAlertRspRemediation struct {

	// The name, display only
	Recommendation *string `json:"recommendation,omitempty"`

	// The name, display only
	Url *string `json:"url,omitempty"`
}

func (o ShowAlertRspRemediation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlertRspRemediation struct{}"
	}

	return strings.Join([]string{"ShowAlertRspRemediation", string(data)}, " ")
}
