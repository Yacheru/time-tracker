package entities

type Task struct {
	ID        int  `json:"id" db:"id"`
	PeopleID  int  `json:"people_id" db:"people_id" binding:"required"`
	StartTask int  `json:"start_task" db:"start_task" binding:"required"`
	EndTask   *int `json:"end_task,omitempty" db:"end_task"`
	Labor     *int `json:"labor,omitempty" db:"labor"`
}
