package main

import "go_todo/app/models"

func main() {
	/*
		fmt.Println(config.Config.Port)
		fmt.Println(config.Config.SQLDriver)
		fmt.Println(config.Config.DbName)
		fmt.Println(config.Config.LogFile)

		log.Println("test")
	*/
	// fmt.Println(models.Db)

	// u := &models.User{}
	// u.Name = "test2"
	// u.Email = "test2@example.com"
	// u.PassWord = "testtest"
	// fmt.Println(u)

	// u.CreateUser()
	// u, _ := models.GetUser(1)
	// fmt.Println(u)

	// // update
	// u.Name = "Test2"
	// u.Email = "test2@email.com"
	// u.UpdateUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// // delete
	// u.DeleteUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// user, _ := models.GetUser(2)
	// user.CreateTodo("First Todo")

	// 取得
	// t, _ := models.GetTodo(1)
	// fmt.Println(t)

	// user, _ := models.GetUser(3)
	// user.CreateTodo("Third Todo")

	// todos, _ := models.GetTodos()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	// user2, _ := models.GetUser(3)
	// todos, _ := user2.GetTodosByUser()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	// t.Content = "Update Todo!"
	// t.UpdateTodo()
	t, _ := models.GetTodo(3)
	t.DeleteTodo()
}
