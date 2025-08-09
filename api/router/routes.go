package router

func InitRoutes(controllers Controllers, r *Router) {

	userRoutesGroup := r.Group("/user")

	userRoutesGroup.Handle("POST", "/create-user", controllers.UserController.CreateUser)

}
