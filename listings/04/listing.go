package main
 
/*
Вывод:
0
...
9
deadlock

Запущенная горутина с циклом завершит свою работу по отправке данных в канал, когда основная горутина (main) продолжит ожидать данные в канале, что вызовет deadlock
*/
func main() {
    ch := make(chan int)
    go func() {
        for i := 0; i < 10; i++ {
            ch <- i
        }
    }()
 
    for n := range ch {
        println(n)
    }
}
