package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNotificationRecordsResponse Response Object
type ListNotificationRecordsResponse struct {

	// Record对象。
	Records *[]Record `json:"records,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListNotificationRecordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNotificationRecordsResponse struct{}"
	}

	return strings.Join([]string{"ListNotificationRecordsResponse", string(data)}, " ")
}
