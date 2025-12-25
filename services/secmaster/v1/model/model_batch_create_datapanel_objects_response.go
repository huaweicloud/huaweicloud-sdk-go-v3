package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateDatapanelObjectsResponse Response Object
type BatchCreateDatapanelObjectsResponse struct {

	// 数据对象信息列表
	Body *[]DataObjectCreateUpdateResponse `json:"body,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchCreateDatapanelObjectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateDatapanelObjectsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateDatapanelObjectsResponse", string(data)}, " ")
}
