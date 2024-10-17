package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/res/v1/model"
)

type CreateResDatasourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateResDatasourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateResDatasourceInvoker) Invoke() (*model.CreateResDatasourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateResDatasourceResponse), nil
	}
}

type CreateResIntelligentSceneInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateResIntelligentSceneInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateResIntelligentSceneInvoker) Invoke() (*model.CreateResIntelligentSceneResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateResIntelligentSceneResponse), nil
	}
}

type CreateResJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateResJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateResJobInvoker) Invoke() (*model.CreateResJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateResJobResponse), nil
	}
}

type CreateResJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateResJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateResJobsInvoker) Invoke() (*model.CreateResJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateResJobsResponse), nil
	}
}

type CreateResOnlineInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateResOnlineInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateResOnlineInstanceInvoker) Invoke() (*model.CreateResOnlineInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateResOnlineInstanceResponse), nil
	}
}

type CreateResSceneInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateResSceneInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateResSceneInvoker) Invoke() (*model.CreateResSceneResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateResSceneResponse), nil
	}
}

type CreateResWorkspaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateResWorkspaceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateResWorkspaceInvoker) Invoke() (*model.CreateResWorkspaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateResWorkspaceResponse), nil
	}
}

type DeleteResDatasourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteResDatasourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteResDatasourceInvoker) Invoke() (*model.DeleteResDatasourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteResDatasourceResponse), nil
	}
}

type DeleteResJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteResJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteResJobInvoker) Invoke() (*model.DeleteResJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteResJobResponse), nil
	}
}

type DeleteResOnlineInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteResOnlineInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteResOnlineInstanceInvoker) Invoke() (*model.DeleteResOnlineInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteResOnlineInstanceResponse), nil
	}
}

type DeleteResSceneInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteResSceneInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteResSceneInvoker) Invoke() (*model.DeleteResSceneResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteResSceneResponse), nil
	}
}

type DeleteResWorkspaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteResWorkspaceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteResWorkspaceInvoker) Invoke() (*model.DeleteResWorkspaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteResWorkspaceResponse), nil
	}
}

type ListResDatasourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResDatasourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListResDatasourcesInvoker) Invoke() (*model.ListResDatasourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResDatasourcesResponse), nil
	}
}

type ListResEnterprisesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResEnterprisesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListResEnterprisesInvoker) Invoke() (*model.ListResEnterprisesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResEnterprisesResponse), nil
	}
}

type ListResOnlineServiceDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResOnlineServiceDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListResOnlineServiceDetailsInvoker) Invoke() (*model.ListResOnlineServiceDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResOnlineServiceDetailsResponse), nil
	}
}

type ListResResourceSpecInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResResourceSpecInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListResResourceSpecInvoker) Invoke() (*model.ListResResourceSpecResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResResourceSpecResponse), nil
	}
}

type ListResScenesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResScenesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListResScenesInvoker) Invoke() (*model.ListResScenesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResScenesResponse), nil
	}
}

type ListResWorkspacesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResWorkspacesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListResWorkspacesInvoker) Invoke() (*model.ListResWorkspacesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResWorkspacesResponse), nil
	}
}

type ShowResDatasourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowResDatasourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowResDatasourceInvoker) Invoke() (*model.ShowResDatasourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResDatasourceResponse), nil
	}
}

type ShowResDatasourceWorkDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowResDatasourceWorkDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowResDatasourceWorkDetailInvoker) Invoke() (*model.ShowResDatasourceWorkDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResDatasourceWorkDetailResponse), nil
	}
}

type ShowResJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowResJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowResJobInvoker) Invoke() (*model.ShowResJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResJobResponse), nil
	}
}

type ShowResRecallSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowResRecallSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowResRecallSetInvoker) Invoke() (*model.ShowResRecallSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResRecallSetResponse), nil
	}
}

type ShowResSceneInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowResSceneInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowResSceneInvoker) Invoke() (*model.ShowResSceneResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResSceneResponse), nil
	}
}

type ShowResWrokspaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowResWrokspaceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowResWrokspaceInvoker) Invoke() (*model.ShowResWrokspaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResWrokspaceResponse), nil
	}
}

type StartResJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartResJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartResJobInvoker) Invoke() (*model.StartResJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartResJobResponse), nil
	}
}

type StartResSceneJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartResSceneJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartResSceneJobsInvoker) Invoke() (*model.StartResSceneJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartResSceneJobsResponse), nil
	}
}

type UpdateResDatasourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateResDatasourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateResDatasourceInvoker) Invoke() (*model.UpdateResDatasourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateResDatasourceResponse), nil
	}
}

type UpdateResDatastructInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateResDatastructInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateResDatastructInvoker) Invoke() (*model.UpdateResDatastructResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateResDatastructResponse), nil
	}
}

type UpdateResIntelligentSceneInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateResIntelligentSceneInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateResIntelligentSceneInvoker) Invoke() (*model.UpdateResIntelligentSceneResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateResIntelligentSceneResponse), nil
	}
}

type UpdateResJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateResJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateResJobInvoker) Invoke() (*model.UpdateResJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateResJobResponse), nil
	}
}

type UpdateResOnlineInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateResOnlineInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateResOnlineInstanceInvoker) Invoke() (*model.UpdateResOnlineInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateResOnlineInstanceResponse), nil
	}
}

type UpdateResSceneInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateResSceneInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateResSceneInvoker) Invoke() (*model.UpdateResSceneResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateResSceneResponse), nil
	}
}

type UpdateResWorkspaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateResWorkspaceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateResWorkspaceInvoker) Invoke() (*model.UpdateResWorkspaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateResWorkspaceResponse), nil
	}
}
