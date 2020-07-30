/*
    * Classroom
    *
    * devcloud classedge api
    *
*/

package model

// Response Object
type ShowJobExercisesResponse struct {
	// 作业下习题列表
	GroupExercises []ExerciseGroup `json:"group_exercises,omitempty"`
	// 作业下习题总数
	Total int32 `json:"total,omitempty"`
}
