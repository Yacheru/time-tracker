package entities

type People struct {
	ID             *int    `json:"id" db:"id"`
	Surname        string  `json:"surname" db:"surname" binding:"required"`
	Name           string  `json:"name" db:"name" binding:"required"`
	Patronymic     *string `json:"patronymic,omitempty" db:"patronymic"`
	PassportSeries int     `json:"passport_series" db:"passport_series" binding:"required"`
	PassportNumber int     `json:"passport_number" db:"passport_number" binding:"required"`
	TaskID         *int    `json:"task_id,omitempty" db:"task_id"`
}
