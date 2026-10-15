package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowObjectColumnInfoResponse Response Object
type ShowObjectColumnInfoResponse struct {

	// 和列信息相关的对象
	ObjectWithColumnInfos *[]ObjectWithColumnInfo `json:"object_with_column_infos,omitempty"`
	HttpStatusCode        int                     `json:"-"`
}

func (o ShowObjectColumnInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowObjectColumnInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowObjectColumnInfoResponse", string(data)}, " ")
}
