package main
 
/*
Вывод:
error

Структура customError имплементирует интерфейс error
В функции test() мы возвращаем nil и присваиваем его интерфейсному типу error
Интерфейс будет nil в том случае, если и динамический (data) и статический (itab) типы равны nil
В нашем случае хоть переменная err в error не указывает ни на какие данные, в ней все еще хранится поле itab, потому err != nil
*/

func main() {
    var err error
    err = test()
    if err != nil {
        println("error")
        return
    }
    println("ok")
}

type customError struct {
	msg string
}

func (e *customError) Error() string {
   return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}
