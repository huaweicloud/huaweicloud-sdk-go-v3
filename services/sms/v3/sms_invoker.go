package v3

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/sms/v3/model"
)

type CreateMigprojectInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateMigprojectInvoker) Invoke() (*model.CreateMigprojectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateMigprojectResponse), nil
	}
}

type CreateTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTaskInvoker) Invoke() (*model.CreateTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTaskResponse), nil
	}
}

type CreateTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTemplateInvoker) Invoke() (*model.CreateTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTemplateResponse), nil
	}
}

type DeleteMigprojectInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteMigprojectInvoker) Invoke() (*model.DeleteMigprojectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteMigprojectResponse), nil
	}
}

type DeleteServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteServerInvoker) Invoke() (*model.DeleteServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteServerResponse), nil
	}
}

type DeleteServersInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteServersInvoker) Invoke() (*model.DeleteServersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteServersResponse), nil
	}
}

type DeleteTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTaskInvoker) Invoke() (*model.DeleteTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTaskResponse), nil
	}
}

type DeleteTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTasksInvoker) Invoke() (*model.DeleteTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTasksResponse), nil
	}
}

type DeleteTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTemplateInvoker) Invoke() (*model.DeleteTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTemplateResponse), nil
	}
}

type DeleteTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTemplatesInvoker) Invoke() (*model.DeleteTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTemplatesResponse), nil
	}
}

type ListErrorServersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListErrorServersInvoker) Invoke() (*model.ListErrorServersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListErrorServersResponse), nil
	}
}

type ListMigprojectsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMigprojectsInvoker) Invoke() (*model.ListMigprojectsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMigprojectsResponse), nil
	}
}

type ListServersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListServersInvoker) Invoke() (*model.ListServersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListServersResponse), nil
	}
}

type ListTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTasksInvoker) Invoke() (*model.ListTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTasksResponse), nil
	}
}

type ListTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTemplatesInvoker) Invoke() (*model.ListTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTemplatesResponse), nil
	}
}

type RegisterServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *RegisterServerInvoker) Invoke() (*model.RegisterServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RegisterServerResponse), nil
	}
}

type ShowCommandInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCommandInvoker) Invoke() (*model.ShowCommandResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCommandResponse), nil
	}
}

type ShowMigprojectInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMigprojectInvoker) Invoke() (*model.ShowMigprojectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMigprojectResponse), nil
	}
}

type ShowOverviewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowOverviewInvoker) Invoke() (*model.ShowOverviewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowOverviewResponse), nil
	}
}

type ShowServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowServerInvoker) Invoke() (*model.ShowServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowServerResponse), nil
	}
}

type ShowTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTaskInvoker) Invoke() (*model.ShowTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTaskResponse), nil
	}
}

type ShowTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTemplateInvoker) Invoke() (*model.ShowTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTemplateResponse), nil
	}
}

type ShowsSpeedLimitsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowsSpeedLimitsInvoker) Invoke() (*model.ShowsSpeedLimitsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowsSpeedLimitsResponse), nil
	}
}

type UpdateCommandResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateCommandResultInvoker) Invoke() (*model.UpdateCommandResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateCommandResultResponse), nil
	}
}

type UpdateCopyStateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateCopyStateInvoker) Invoke() (*model.UpdateCopyStateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateCopyStateResponse), nil
	}
}

type UpdateDefaultMigprojectInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDefaultMigprojectInvoker) Invoke() (*model.UpdateDefaultMigprojectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDefaultMigprojectResponse), nil
	}
}

type UpdateDiskInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDiskInfoInvoker) Invoke() (*model.UpdateDiskInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDiskInfoResponse), nil
	}
}

type UpdateMigprojectInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMigprojectInvoker) Invoke() (*model.UpdateMigprojectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMigprojectResponse), nil
	}
}

type UpdateServerNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateServerNameInvoker) Invoke() (*model.UpdateServerNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateServerNameResponse), nil
	}
}

type UpdateSpeedInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSpeedInvoker) Invoke() (*model.UpdateSpeedResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSpeedResponse), nil
	}
}

type UpdateTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTaskInvoker) Invoke() (*model.UpdateTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTaskResponse), nil
	}
}

type UpdateTaskSpeedInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTaskSpeedInvoker) Invoke() (*model.UpdateTaskSpeedResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTaskSpeedResponse), nil
	}
}

type UpdateTaskStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTaskStatusInvoker) Invoke() (*model.UpdateTaskStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTaskStatusResponse), nil
	}
}

type UpdateTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTemplateInvoker) Invoke() (*model.UpdateTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTemplateResponse), nil
	}
}
