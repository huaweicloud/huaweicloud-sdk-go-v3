package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublishVersionVoSearchResultData 返回数据。
type PublishVersionVoSearchResultData struct {
	Value *PublishVersionVoSearchResultDataValue `json:"value,omitempty"`
}

func (o PublishVersionVoSearchResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishVersionVoSearchResultData struct{}"
	}

	return strings.Join([]string{"PublishVersionVoSearchResultData", string(data)}, " ")
}
