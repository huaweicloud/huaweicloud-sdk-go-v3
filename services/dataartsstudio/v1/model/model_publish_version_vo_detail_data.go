package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublishVersionVoDetailData 接口返回的数据。
type PublishVersionVoDetailData struct {
	Value *PublishVersionVo `json:"value,omitempty"`
}

func (o PublishVersionVoDetailData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishVersionVoDetailData struct{}"
	}

	return strings.Join([]string{"PublishVersionVoDetailData", string(data)}, " ")
}
