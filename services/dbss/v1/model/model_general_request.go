package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GeneralRequest struct {

	// 数据库ID
	DbId *string `json:"db_id,omitempty"`

	// 是否重新生成  - 1：是  - 0：否
	Regenerate int32 `json:"regenerate"`

	Time *Duration `json:"time"`
}

func (o GeneralRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GeneralRequest struct{}"
	}

	return strings.Join([]string{"GeneralRequest", string(data)}, " ")
}
