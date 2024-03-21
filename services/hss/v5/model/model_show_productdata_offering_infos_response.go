package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProductdataOfferingInfosResponse Response Object
type ShowProductdataOfferingInfosResponse struct {

	//
	Body           *[]ResourceProductDataObjectInfo `json:"body,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o ShowProductdataOfferingInfosResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProductdataOfferingInfosResponse struct{}"
	}

	return strings.Join([]string{"ShowProductdataOfferingInfosResponse", string(data)}, " ")
}
