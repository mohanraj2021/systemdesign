package main

type User struct {
	Name     string
	Password string
}

type NotificationService interface {
	SendNotification(user User, message string)
}

type UserRepo interface {
	AddUser(user User)
	GetUser(name string) (User, bool)
}

type UserRepository struct {
	Users map[string]User
}

func (r *UserRepository) AddUser(user User) {
	r.Users[user.Name] = user
}

func (r *UserRepository) GetUser(name string) (User, bool) {
	if user, ok := r.Users[name]; ok {
		return user, true
	}
	return User{}, false
}

func NewRepository() UserRepo {
	return &UserRepository{
		Users: make(map[string]User, 0),
	}
}

func (r *AuthenticationService) ChangeNotificationService(mailService NotificationService) {
	r.MailService = mailService
}

type AuthenticationService struct {
	userRepo    UserRepo
	MailService NotificationService
}

type EmailNotificationService struct{}

func (e *EmailNotificationService) SendNotification(user User, message string) {
	println("Sending email to", user.Name, "with message:", message)
}

type SMSNotificationService struct{}

func (s *SMSNotificationService) SendNotification(user User, message string) {
	println("Sending SMS to", user.Name, "with message:", message)
}

func NewAuthenticationService(userRepo UserRepo, mailService NotificationService) AuthenticationService {
	return AuthenticationService{userRepo: userRepo, MailService: mailService}
}

func main() {

	emailService := &EmailNotificationService{}
	repo := NewRepository()
	service := NewAuthenticationService(repo, emailService)

	service.userRepo.AddUser(User{Name: "Alice", Password: "password123"})
	user, t := service.userRepo.GetUser("Alice")
	if !t {
		println("No user found in repository")
		return
	}
	service.MailService.SendNotification(user, "Welcome to our service!")
	println("Retrieved User:", user.Name)

	smsService := &SMSNotificationService{}

	service.ChangeNotificationService(smsService)
	service.MailService.SendNotification(user, "Your package has been shipped")

}
