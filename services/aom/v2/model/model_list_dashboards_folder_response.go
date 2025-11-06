package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDashboardsFolderResponse Response Object
type ListDashboardsFolderResponse struct {
	Body           *[]DashBoardsFolder `json:"body,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListDashboardsFolderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDashboardsFolderResponse struct{}"
	}

	return strings.Join([]string{"ListDashboardsFolderResponse", string(data)}, " ")
}
