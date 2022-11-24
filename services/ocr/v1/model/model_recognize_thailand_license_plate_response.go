package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RecognizeThailandLicensePlateResponse struct {
	Result         *ThailandLicensePlateItem `json:"result,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o RecognizeThailandLicensePlateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeThailandLicensePlateResponse struct{}"
	}

	return strings.Join([]string{"RecognizeThailandLicensePlateResponse", string(data)}, " ")
}
