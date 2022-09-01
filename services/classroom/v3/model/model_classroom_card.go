package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ClassroomCard struct {

	// 课堂ID
	ClassroomId string `json:"classroom_id" xml:"classroom_id"`

	// 课堂名称
	Name string `json:"name" xml:"name"`

	// 课堂描述
	Description string `json:"description" xml:"description"`

	// 课堂学分
	Credit float32 `json:"credit" xml:"credit"`

	// 课堂当前的状态，normal：课堂处于正常状态，archive：课堂已归档
	Status string `json:"status" xml:"status"`
}

func (o ClassroomCard) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClassroomCard struct{}"
	}

	return strings.Join([]string{"ClassroomCard", string(data)}, " ")
}
