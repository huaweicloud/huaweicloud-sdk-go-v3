package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubscriptionProductResponse Response Object
type ListSubscriptionProductResponse struct {
	Basic *ProductRspBasic `json:"basic,omitempty"`

	Standard *ProductRspStandard `json:"standard,omitempty"`

	Professional *ProductRspProfessional `json:"professional,omitempty"`

	LargeScreen *ProductRspLargeScreen `json:"large_screen,omitempty"`

	LogCollection *ProductRspLogCollection `json:"log_collection,omitempty"`

	LogRetention *ProductRspLogRetention `json:"log_retention,omitempty"`

	LogAnalysis *ProductRspLogAnalysis `json:"log_analysis,omitempty"`

	Soar           *ProductRspSoar `json:"soar,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListSubscriptionProductResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubscriptionProductResponse struct{}"
	}

	return strings.Join([]string{"ListSubscriptionProductResponse", string(data)}, " ")
}
