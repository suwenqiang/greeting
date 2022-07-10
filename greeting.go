package greeting

// add tag v0.1.0
import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return name, errors.New("empty name")
	}
	message := fmt.Sprintf(randFormate(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
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
