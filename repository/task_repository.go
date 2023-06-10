// /repository/task_repository.go

package repository

import (
	"fmt"
	"go-echo-api-sample-202306/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// IFの定義
type ITaskRepository interface {
	GetAllTasks(tasks *[]model.Task) error
	GetTaskById(task *model.Task, taskId uint) error
	CreateTask(task *model.Task) error
	UpdateTask(task *model.Task, taskId uint) error
	DeleteTask(taskId uint) error
}

type taskRepository struct {
	db *gorm.DB
}

// コンストラクタの作成、main.goで呼び出し
func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &taskRepository{db}
}

// タスク全取得
func (tr *taskRepository) GetAllTasks(tasks *[]model.Task) error {
	if err := tr.db.Order("created_at").Find(tasks).Error; err != nil {
		return err
	}
	return nil
}

// タスク取得
func (tr *taskRepository) GetTaskById(task *model.Task, taskId uint) error {
	if err := tr.db.First(task, taskId).Error; err != nil {
		return err
	}
	return nil
}

// タスク追加
func (tr *taskRepository) CreateTask(task *model.Task) error {
	if err := tr.db.Create(task).Error; err != nil {
		return err
	}
	return nil
}

// タスク更新
func (tr *taskRepository) UpdateTask(task *model.Task, taskId uint) error {
	result := tr.db.Model(task).Clauses(clause.Returning{}).Where("id=?", taskId).Update("title", task.Title)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

// タスク削除
func (tr *taskRepository) DeleteTask(taskId uint) error {
	result := tr.db.Where("id=?", taskId).Delete(&model.Task{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
