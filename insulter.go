package insulter

import (
	"math/rand"
	"time"
	)


func pickWord(x []string) string {
	rand.Seed(time.Now().UnixNano())
	return x[rand.Intn(len(x))]
}

func Insult(name string) string {
	list := []string{
		"Красавчик",
		"Идиот",
		"Какашка",
		"Гриб",
		"Енот",
		"Победитель",
		"Хороший человек",
		"Обычный, ничем не примечательный чел",
	}
	swear:= pickWord(list)
	answer:= "Кто сегодня  "+ name+ "? "+name+" сегодня "+ swear+ "!"
return answer
}
