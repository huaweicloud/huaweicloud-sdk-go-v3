package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDashboardsFolderResponse Response Object
type DeleteDashboardsFolderResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDashboardsFolderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDashboardsFolderResponse struct{}"
	}

	return strings.Join([]string{"DeleteDashboardsFolderResponse", string(data)}, " ")
}
