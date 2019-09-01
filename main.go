package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

func main() {
	logSimpleDevelopment()
	logSugarDevelopment()
	logSimpleProduction()
	logSugarProduction()
	logSugarExample()
	logObject()
}

func logSimpleDevelopment() {
	fmt.Println("== log simple development start ==")
	logger, _ := zap.NewDevelopment()
	logger.Info("Hello World!", zap.String("key", "value"), zap.Time("now", time.Now()))
	logger.With(zap.Namespace("test namespace")).Info("Hello World!", zap.String("key", "value"), zap.Time("now", time.Now()))
	fmt.Println("== log simple development end ==")
}

func logSugarDevelopment() {
	fmt.Println("== log sugar development start ==")
	logger, _ := zap.NewDevelopment()
	sugar := logger.Sugar()
	sugar.Info("Hello", "World!")
	sugar.Infof("name: %s, age: %d", "jon", 12)
	sugar.Infow("Hello World!", zap.String("key", "value"), zap.Time("now", time.Now()))
	sugar.With(zap.Namespace("test namespace")).Infow("Hello World!", zap.String("key", "value"), zap.Time("now", time.Now()))
	_, err := os.Open("not_exist.txt")
	sugar.Errorw("File not Exists", zap.Error(err), zap.Time("now", time.Now()))
	fmt.Println("== log sugar development end ==")
}

func logSimpleProduction() {
	fmt.Println("== log simple production start ==")
	logger, _ := zap.NewProduction()
	logger.Info("Hello World!", zap.String("key", "value"), zap.Time("now", time.Now()))
	fmt.Println("== log simple production end ==")
}

func logSugarProduction() {
	fmt.Println("== log sugar production start ==")
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()
	sugar.Info("Hello", "World!")
	sugar.Infof("name: %s, age: %d", "jon", 12)
	sugar.Infow("Hello World!", zap.String("key", "value"), zap.Time("now", time.Now()))
	_, err := os.Open("not_exist.txt")
	sugar.Errorw("File not Exists", zap.Error(err), zap.Time("now", time.Now()))
	fmt.Println("== log sugar production end ==")
}

func logSugarExample() {
	fmt.Println("== log sugar start ==")
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()
	sugar.Infow("failed to fetch URL",
		"url", "http://example.com",
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("failed to fetch URL: %s", "http://example.com")
	fmt.Println("== log sugar end==")
}

func logObject() {
	user := &user{
		Name: "jon",
		Age:  12,
	}
	logger, _ := zap.NewDevelopment()
	logger.Info("object sample", zap.Object("userObj", user))
	logger.With(zap.Namespace("test namespace")).Info("object sample", zap.Object("userObj", user))
	loggerPro, _ := zap.NewProduction()
	loggerPro.Info("object sample", zap.Object("userObj", user))
	loggerPro.With(zap.Namespace("test namespace")).Info("object sample", zap.Object("userObj", user))
}

type user struct {
	Name string
	Age  int64
}

func (u user) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("name", u.Name)
	enc.AddInt64("age", u.Age)
	return nil
}
