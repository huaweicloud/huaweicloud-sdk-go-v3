package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteOneWidgetResponse Response Object
type DeleteOneWidgetResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteOneWidgetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteOneWidgetResponse struct{}"
	}

	return strings.Join([]string{"DeleteOneWidgetResponse", string(data)}, " ")
}
