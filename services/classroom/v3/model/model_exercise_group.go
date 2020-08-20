/*
 * Classroom
 *
 * devcloud classedge api
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

type ExerciseGroup struct {
	// 习题列表
	Exercises []ExerciseCard `json:"exercises"`
	// 习题分类
	Type string `json:"type"`
}

func (o ExerciseGroup) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ExerciseGroup", string(data)}, " ")
}
