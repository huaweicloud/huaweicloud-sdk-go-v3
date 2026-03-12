package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteThumbnailsResponse Response Object
type DeleteThumbnailsResponse struct {

	// 删除截图信息的处理结果。
	DeleteResultArray *[]DeleteThumbnailResult `json:"delete_result_array,omitempty"`
	HttpStatusCode    int                      `json:"-"`
}

func (o DeleteThumbnailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteThumbnailsResponse struct{}"
	}

	return strings.Join([]string{"DeleteThumbnailsResponse", string(data)}, " ")
}
