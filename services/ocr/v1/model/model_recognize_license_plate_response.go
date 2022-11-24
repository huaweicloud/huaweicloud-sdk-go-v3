package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RecognizeLicensePlateResponse struct {
	Result         *LicensePlateResult `json:"result,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o RecognizeLicensePlateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeLicensePlateResponse struct{}"
	}

	return strings.Join([]string{"RecognizeLicensePlateResponse", string(data)}, " ")
}
