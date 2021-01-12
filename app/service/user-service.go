package service

import (
    "log"

    "github.com/mashingan/smapping"
    "golang_api/app/http/request"
    "golang_api/app/model"
    "golang_api/app/repository"
)

type UserService interface {
    Update(user request.UserUpdateRequest) model.User
    Profile(userID string) model.User
}

type userService struct {
    userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
    return &userService{
        userRepository: userRepo,
    }
}

func (service *userService) Update(user request.UserUpdateRequest) model.User {
    userToUpdate := model.User{}
    err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
    if err != nil {
        log.Fatalf("Failed map %v:", err)
    }
    updatedUser := service.userRepository.UpdateUser(userToUpdate)
    return updatedUser
}

func (service *userService) Profile(userID string) model.User {
    return service.userRepository.ProfileUser(userID)
}
