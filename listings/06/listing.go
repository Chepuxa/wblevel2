package main
 
import (
  "fmt"
)
 
/*
Вывод:
3 2 3

В рантайме представлены структурой
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}
Где array - указатель на аллоцированный массив, len - текущая длина среза, cap - вместимость среза

В Go при передаче значения в функцию, эта функция будет операрировать с копией данных. Но т.к. slice - по сути указатель на внутрилежащий массив, функция будет работать с тем же адресом в памяти
Из-за чего изменение элемента 0 на строке 32 отобразиться также во всех экземплярах этого слайса
Встроенная функция append добавляет элемент в конец среза.
Если cap не достаточно, то будет выделено новое хранилище (будет указывать на новый массив).
На строке 33 после добавления элемента в слайс переменная i будет указывать уже на новый массив (т.к. мы будем увеличивать cap слайса), и дальнейшие его изменения не отобразятся в изначальном слайсе
*/
func main() {
  var s = []string{"1", "2", "3"}
  modifySlice(s)
  fmt.Println(s)
}
 
func modifySlice(i []string) {
  i[0] = "3"
  i = append(i, "4")
  i[1] = "5"
  i = append(i, "6")
}