package service

import (
    "fmt"
    "log"

    "github.com/mashingan/smapping"
    "golang_api/app/http/request"
    "golang_api/app/model"
    "golang_api/app/repository"
)

type BookService interface {
    Insert(b request.BookCreateRequest) model.Book
    Update(b request.BookUpdateRequest) model.Book
    Delete(b model.Book)
    All() []model.Book
    FindByID(bookID uint64) model.Book
    IsAllowedToEdit(userID string, bookID uint64) bool
}

type bookService struct {
    bookRepository repository.BookRepository
}

func NewBookService(bookRepo repository.BookRepository) BookService {
    return &bookService{
        bookRepository: bookRepo,
    }
}

func (service *bookService) Insert(b request.BookCreateRequest) model.Book {
    book := model.Book{}
    err := smapping.FillStruct(&book, smapping.MapFields(&b))
    if err != nil {
        log.Fatalf("Failed map %v: ", err)
    }
    res := service.bookRepository.InsertBook(book)
    return res
}

func (service *bookService) Update(b request.BookUpdateRequest) model.Book {
    book := model.Book{}
    err := smapping.FillStruct(&book, smapping.MapFields(&b))
    if err != nil {
        log.Fatalf("Failed map %v: ", err)
    }
    res := service.bookRepository.UpdateBook(book)
    return res
}

func (service *bookService) Delete(b model.Book) {
    service.bookRepository.DeleteBook(b)
}

func (service *bookService) All() []model.Book {
    return service.bookRepository.AllBook()
}

func (service *bookService) FindByID(bookID uint64) model.Book {
    return service.bookRepository.FindBookByID(bookID)
}

func (service *bookService) IsAllowedToEdit(userID string, bookID uint64) bool {
    b := service.bookRepository.FindBookByID(bookID)
    id := fmt.Sprintf("%v", b.UserID)
    return userID == id
}
