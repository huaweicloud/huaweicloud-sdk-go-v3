package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDriftDetailsResponse Response Object
type ListDriftDetailsResponse struct {

	// 漂移详细信息。
	DriftDetails   *[]DriftDetail `json:"drift_details,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListDriftDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDriftDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListDriftDetailsResponse", string(data)}, " ")
}
