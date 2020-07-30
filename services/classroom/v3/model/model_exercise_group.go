/*
    * Classroom
    *
    * devcloud classedge api
    *
*/

package model

type ExerciseGroup struct {
	// 习题列表
	Exercises []ExerciseCard `json:"exercises"`
	// 习题分类
	Type string `json:"type"`
}
