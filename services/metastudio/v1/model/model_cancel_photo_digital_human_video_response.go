package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelPhotoDigitalHumanVideoResponse Response Object
type CancelPhotoDigitalHumanVideoResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CancelPhotoDigitalHumanVideoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelPhotoDigitalHumanVideoResponse struct{}"
	}

	return strings.Join([]string{"CancelPhotoDigitalHumanVideoResponse", string(data)}, " ")
}
