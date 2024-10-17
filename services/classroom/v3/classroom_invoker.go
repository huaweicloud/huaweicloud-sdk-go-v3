package v3

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/classroom/v3/model"
)

type ApplyJudgementInvoker struct {
	*invoker.BaseInvoker
}

func (i *ApplyJudgementInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ApplyJudgementInvoker) Invoke() (*model.ApplyJudgementResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ApplyJudgementResponse), nil
	}
}

type ShowJudgementDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJudgementDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowJudgementDetailInvoker) Invoke() (*model.ShowJudgementDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJudgementDetailResponse), nil
	}
}

type ShowJudgementFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJudgementFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowJudgementFileInvoker) Invoke() (*model.ShowJudgementFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJudgementFileResponse), nil
	}
}

type ExecuteExerciseInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteExerciseInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteExerciseInvoker) Invoke() (*model.ExecuteExerciseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteExerciseResponse), nil
	}
}

type ListExercisesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListExercisesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListExercisesInvoker) Invoke() (*model.ListExercisesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListExercisesResponse), nil
	}
}

type ListPackagesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPackagesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPackagesInvoker) Invoke() (*model.ListPackagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPackagesResponse), nil
	}
}

type ShowExerciseDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowExerciseDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowExerciseDetailInvoker) Invoke() (*model.ShowExerciseDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowExerciseDetailResponse), nil
	}
}

type ShowPackageDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPackageDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPackageDetailInvoker) Invoke() (*model.ShowPackageDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPackageDetailResponse), nil
	}
}

type ListAllDifficultsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAllDifficultsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAllDifficultsInvoker) Invoke() (*model.ListAllDifficultsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAllDifficultsResponse), nil
	}
}

type ListMyKnowledgePointsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMyKnowledgePointsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMyKnowledgePointsInvoker) Invoke() (*model.ListMyKnowledgePointsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMyKnowledgePointsResponse), nil
	}
}

type ListClassroomMembersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListClassroomMembersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListClassroomMembersInvoker) Invoke() (*model.ListClassroomMembersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListClassroomMembersResponse), nil
	}
}

type ListClassroomsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListClassroomsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListClassroomsInvoker) Invoke() (*model.ListClassroomsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListClassroomsResponse), nil
	}
}

type ShowClassroomDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowClassroomDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowClassroomDetailInvoker) Invoke() (*model.ShowClassroomDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowClassroomDetailResponse), nil
	}
}

type ListClassroomMemberJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListClassroomMemberJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListClassroomMemberJobsInvoker) Invoke() (*model.ListClassroomMemberJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListClassroomMemberJobsResponse), nil
	}
}

type ListJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListJobsInvoker) Invoke() (*model.ListJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListJobsResponse), nil
	}
}

type ListMemberJobRecordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMemberJobRecordsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMemberJobRecordsInvoker) Invoke() (*model.ListMemberJobRecordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMemberJobRecordsResponse), nil
	}
}

type ShowJobDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowJobDetailInvoker) Invoke() (*model.ShowJobDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobDetailResponse), nil
	}
}

type ShowJobExercisesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobExercisesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowJobExercisesInvoker) Invoke() (*model.ShowJobExercisesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobExercisesResponse), nil
	}
}
