package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePeriodToOnDemandResponse Response Object
type UpdatePeriodToOnDemandResponse struct {

	// HTTP 200的时候返回该字段；部分失败时仅返回失败的记录；如果全部成功，则该记录为空，具体参见表1。
	ErrorDetails   *[]ErrorDetail `json:"error_details,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o UpdatePeriodToOnDemandResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePeriodToOnDemandResponse struct{}"
	}

	return strings.Join([]string{"UpdatePeriodToOnDemandResponse", string(data)}, " ")
}
