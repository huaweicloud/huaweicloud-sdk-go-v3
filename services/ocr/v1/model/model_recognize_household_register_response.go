package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecognizeHouseholdRegisterResponse Response Object
type RecognizeHouseholdRegisterResponse struct {

	// 调用成功时表示调用结果。  调用失败时无此字段。
	Result *[]HouseholdRegisterResult `json:"result,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RecognizeHouseholdRegisterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeHouseholdRegisterResponse struct{}"
	}

	return strings.Join([]string{"RecognizeHouseholdRegisterResponse", string(data)}, " ")
}
