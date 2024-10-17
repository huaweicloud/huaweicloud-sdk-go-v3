package v5

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/codecraft/v5/model"
)

type CreateCompetitionScoreInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCompetitionScoreInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCompetitionScoreInvoker) Invoke() (*model.CreateCompetitionScoreResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCompetitionScoreResponse), nil
	}
}

type ListCompetitionWorksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCompetitionWorksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCompetitionWorksInvoker) Invoke() (*model.ListCompetitionWorksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCompetitionWorksResponse), nil
	}
}

type RegisterCompetitionInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *RegisterCompetitionInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RegisterCompetitionInfoInvoker) Invoke() (*model.RegisterCompetitionInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RegisterCompetitionInfoResponse), nil
	}
}

type UpdateCompetitionScoreInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateCompetitionScoreInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateCompetitionScoreInvoker) Invoke() (*model.UpdateCompetitionScoreResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateCompetitionScoreResponse), nil
	}
}
