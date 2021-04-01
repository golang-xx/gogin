package main
import "github.com/gin-gonic/gin"

type UserController struct {
	*gin.Engine
}
// 构造函数
func NewUserController(e *gin.Engine) *UserController {
	return &UserController{e}
}

// 获取用户 业务逻辑
func (this *UserController)GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{
			"data":"hello world",
		})
	}
}

func (this *UserController) Router() {
	this.Handle("GET","/",this.GetUser())
}

func main(){
	r := gin.Default()
	NewUserController(r).Router()
	r.Run(":8091")
}