package main

import (
	"go-echo-api-sample-202306/controller"
	"go-echo-api-sample-202306/db"
	"go-echo-api-sample-202306/repository"
	"go-echo-api-sample-202306/router"
	"go-echo-api-sample-202306/usecase"
	"go-echo-api-sample-202306/validator"
)

func main() {
	db := db.NewDB()
	taskValidator := validator.NewTaskValidator()
	taskRepository := repository.NewTaskRepository(db)                   // repositoryで作ったコンストラクタを起動
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator) // usecaseで作ったコンストラクタを起動
	taskController := controller.NewTaskController(taskUsecase)          // controllerで作ったコンストラクタを起動
	e := router.NewRouter(taskController)                                // routerで作ったコンストラクタを起動
	e.Logger.Fatal(e.Start(":8080"))                                     // e.Startでサーバ起動、エラーの場合、ログ情報を出力し強制終了
}
