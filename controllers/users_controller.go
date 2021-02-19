package controllers

var (
	UserController userControllerInterface = &userController{}
)

type userControllerInterface interface {
	Create()
}
type userController struct {
}

func (c *userController) Create() {

}
