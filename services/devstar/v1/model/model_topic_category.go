package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TopicCategory struct {

	// topic的id。
	TopicId *string `json:"topic_id,omitempty" xml:"topic_id"`

	// topic的名称。
	TopicName *string `json:"topic_name,omitempty" xml:"topic_name"`

	// topic对应的类别的id。
	CategoryId *string `json:"category_id,omitempty" xml:"category_id"`

	// topic对应的类别的名称。
	CategoryName *string `json:"category_name,omitempty" xml:"category_name"`
}

func (o TopicCategory) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopicCategory struct{}"
	}

	return strings.Join([]string{"TopicCategory", string(data)}, " ")
}
