package v3

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/classroom/v3/model"
)

type ApplyJudgementInvoker struct {
	*invoker.BaseInvoker
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

func (i *ShowJudgementFileInvoker) Invoke() (*model.ShowJudgementFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJudgementFileResponse), nil
	}
}

type ListClassroomMembersInvoker struct {
	*invoker.BaseInvoker
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

func (i *ShowJobExercisesInvoker) Invoke() (*model.ShowJobExercisesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobExercisesResponse), nil
	}
}
