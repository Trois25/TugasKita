package entity

type TaskDataInterface interface {
	CreateTask(input TaskCore) error
	FindAllTask() ([]TaskCore, error)
	FindById(taskId string) (TaskCore, error)
	UpdateTask(taskId string, data TaskCore) error
	DeleteTask(taskId string) error
	
	UpdateTaskStatus(taskId string, data UserTaskUploadCore) error
	FindUserTaskById(id string) (UserTaskUploadCore, error)
	FindAllUserTask()([]UserTaskUploadCore, error)

	UploadTask(input UserTaskUploadCore) error
	UploadTaskRequest(input UserTaskSubmissionCore) error
	FindAllRequestTask()([]UserTaskSubmissionCore, error)
	FindAllClaimedTask(userId string) ([]UserTaskUploadCore, error)
	FindAllRequestTaskHistory(userId string)([]UserTaskSubmissionCore, error)
	FindTasksNotClaimedByUser(userId string) ([]TaskCore, error)
}

type TaskUseCaseInterface interface {
	CreateTask(input TaskCore) error
	FindAllTask() ([]TaskCore, error)
	FindById(taskId string) (TaskCore, error)
	UpdateTask(taskId string, data TaskCore) error
	DeleteTask(taskId string) error
	
	UpdateTaskStatus(taskId string, data UserTaskUploadCore) error
	FindUserTaskById(id string) (UserTaskUploadCore, error)
	FindAllUserTask()([]UserTaskUploadCore, error)

	UploadTask(input UserTaskUploadCore) error
	UploadTaskRequest(input UserTaskSubmissionCore) error
	FindAllRequestTask()([]UserTaskSubmissionCore, error)
	FindAllClaimedTask(userId string) ([]UserTaskUploadCore, error)
	FindAllRequestTaskHistory(userId string)([]UserTaskSubmissionCore, error)
	FindTasksNotClaimedByUser(userId string) ([]TaskCore, error)
}
