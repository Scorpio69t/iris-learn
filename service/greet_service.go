package service

import (
	"iris-learn/database"
	"iris-learn/environment"
	"log"
)

// GreetService 是打招呼服务接口
type GreetService interface {
	Say(input string) (string, error)
}

// NewGreetService 创建 GreetService 实例
func NewGreetService(env environment.Env, db database.DB) GreetService {
	service := &greeter{db: db, prefix: "hello"}

	err := service.db.Init("greet.db")
	if err != nil {
		panic(err)
	}

	switch env {
	case environment.PROD:
		return service
	case environment.DEV:
		return &greeterWithLogging{service}
	default:
		panic("unknown environment")
	}
}

// greeter 是打招呼服务
type greeter struct {
	prefix string
	db     database.DB
}

// Say 打招呼
func (s *greeter) Say(input string) (string, error) {
	return s.prefix + " " + input, nil
}

// greeterWithLogging 是带日志的打招呼服务
type greeterWithLogging struct {
	*greeter
}

// Say 打招呼
func (s *greeterWithLogging) Say(input string) (string, error) {
	log.Printf("Say: input=%s", input)
	return s.greeter.Say(input)
}
