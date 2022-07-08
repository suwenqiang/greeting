package greeting

import "fmt"
import "errors"
import "math/rand"
import "time"

func Hello(name string) (string, error) {
	if name == "" {
		return name, errors.New("empty name")
	}
	message := fmt.Sprintf(randFormate(), name)
	return message, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randFormate() string {
	formates := []string{
		"Hi, %v. Welcome",
		"Great to see you, %v",
		"Hail, %v! Well met!",
	}
	return formates[rand.Intn(len(formates))]
}
