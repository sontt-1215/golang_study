package routes

import (
    "golang_api/app/http/controller"
    "golang_api/app/http/middleware"
    "golang_api/app/service"
    "golang_api/app/repository"
    "github.com/gin-gonic/gin"
    "golang_api/config"
    "gorm.io/gorm"
)

var (
    db             *gorm.DB                  = config.SetupDatabaseConnection()
    userRepository repository.UserRepository = repository.NewUserRepository(db)
    bookRepository repository.BookRepository = repository.NewBookRepository(db)
    jwtService     service.JWTService        = service.NewJWTService()
    userService    service.UserService       = service.NewUserService(userRepository)
    bookService    service.BookService       = service.NewBookService(bookRepository)
    authService    service.AuthService       = service.NewAuthService(userRepository)
    authController controller.AuthController = controller.NewAuthController(authService, jwtService)
    userController controller.UserController = controller.NewUserController(userService, jwtService)
    bookController controller.BookController = controller.NewBookController(bookService, jwtService)
)

func InitRouter() *gin.Engine {
    routes := gin.Default()

    authRoutes := routes.Group("api/auth")
    {
        authRoutes.POST("/login", authController.Login)
        authRoutes.POST("/register", authController.Register)
    }

    userRoutes := routes.Group("api/user", middleware.AuthorizeJWT(jwtService))
    {
        userRoutes.GET("/profile", userController.Profile)
        userRoutes.PUT("/profile", userController.Update)
    }

    bookRoutes := routes.Group("api/books", middleware.AuthorizeJWT(jwtService))
    {
        bookRoutes.GET("/", bookController.All)
        bookRoutes.POST("/", bookController.Insert)
        bookRoutes.GET("/:id", bookController.FindByID)
        bookRoutes.PUT("/:id", bookController.Update)
        bookRoutes.DELETE("/:id", bookController.Delete)
    }

    return routes
}
