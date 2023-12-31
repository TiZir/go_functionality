# Устные вопросы
## 1. Какой самый эффективный способ конкатенации строк?
- Go для конкатенации строк обычно используются два метода из пакета strings.
- Пакет strings предоставляет функцию Join, которая может быть использована для конкатенации среза строк. Этот подход эффективен, так как функция Join объединяет строки без создания лишних промежуточных строк.
```go
	strSlice := []string{"Hello,", " ", "World!"}
	str := strings.Join(strSlice, "")
```
- Если нужно конкатенировать несколько строк в цикле или в функции с большим количеством операций конкатенации, то использование strings.Builder может быть более эффективным. strings.Builder предоставляет буферизированную строку, в которую можно эффективно добавлять текст.
```go
	var builder strings.Builder
	builder.WriteString("Hello, ")
	builder.WriteString("World!")
	result := builder.String()
	fmt.Println(result)
```

## 2. Что такое интерфейсы, как они применяются в Go?
- В Go интерфейсы являются набором методов, которые описывают поведение определенного типа данных. Они определяют сигнатуры методов, но не содержат их реализацию. Интерфейсы позволяют группировать типы данных по их общим способностям и позволяют взаимодействовать с объектами, не привязываясь к конкретному типу, так как любой тип в Go соответствует пустому интерфейсу ```interface{}```.
```go
    package main

    import (
        "fmt"
    )

    // Определение интерфейса
    type Shape interface {
        Area() float64
    }

    // Определение структуры, реализующей интерфейс Shape
    type Circle struct {
        Radius float64
    }

    // Реализация метода Area() для структуры Circle
    func (c Circle) Area() float64 {
        return 3.14 * c.Radius * c.Radius
    }

    // Определение структуры, реализующей интерфейс Shape
    type Rectangle struct {
        Width  float64
        Height float64
    }

    // Реализация метода Area() для структуры Rectangle
    func (r Rectangle) Area() float64 {
        return r.Width * r.Height
    }

    // Функция, принимающая любой тип, реализующий интерфейс Shape
    func PrintArea(s Shape) {
        fmt.Printf("Площадь фигуры: %f\n", s.Area())
    }

    func main() {
        circle := Circle{
            Radius: 5,
        }

        rectangle := Rectangle{
            Width:  4,
            Height: 3,
        }

        // Вызов функции с различными типами данных, реализующими интерфейс Shape
        PrintArea(circle)
        PrintArea(rectangle)
    }
```

## 3. Чем отличаются RWMutex от Mutex?

- Mutex используется для обеспечения эксклюзивного доступа к общему ресурсу. Когда горутины получают мьютекс, только одна горутина может получить доступ к ресурсу, пока другие горутины не освободят мьютекс. Это гарантирует, что только один поток будет выполнять критическую секцию кода.

- RWMutex используется для обеспечения параллельного доступа к ресурсу. Он имеет два режима: чтение и запись. Множество горутин может одновременно получать мьютекс для чтения, если никто не пытается получить мьютекс для записи. В противном случае, если горутина попытается получить мьютекс для записи, то она будет заблокирована до тех пор, пока все горутины, которые получили мьютекс на чтение, не освободят его.

- Таким образом, RWMutex предоставляют преимущество в ситуациях, где чтение данных является более распространенной операцией, чем их запись. Это может помочь увеличить параллелизм и производительность программы.

- Также RWMutex показывает чуть лучшую производительность при большом количестве читателей (от 512).

## 4. Чем отличаются буферизированные и не буферизированные каналы?
- Не буферизированные каналы - это каналы, у которых отсутствует внутренний буфер для хранения значений. Когда отправляется значение в не буферизированный канал, отправка будет заблокирована до тех пор, пока получатель не получит значение. Аналогично и наоборот при получении.
- Буферизированные каналы имеют внутренний буфер фиксированного размера. Это позволяет отложить блокировку отправки или получения данных. Когда отправляется значение в буферизированный канал, отправка будет блокирована только если буфер заполнен. Аналогично и наоборот при получении.
- Основное отличие между буферизированными и не буферизированными каналами заключается в блокировке операций отправки и получения данных. В буферизированных каналах блокировка возможна только при полном заполнении или опустошении буфера, в то время как в не буферизированных каналах блокировка происходит всегда до момента отправки и получения данных.

## 5. Какой размер у структуры struct{}{}?
- Размер пустой структуры - 0 байт, так как ее размер определяется суммой размеров типов данных для полей, из которых она состоит. 

## 6. Есть ли в Go перегрузка методов или операторов?
- В Go нет перегрузки методов или операторов.

## 7. В какой последовательности будут выведены элементы map[int]int?
Пример:
```go
    m[0]=1
    m[1]=124
    m[2]=281
```
- Порядок итерации элементов в map в Go является неопределенным и выводятся в случайной последовательности. Это связано с тем, что map реализован как хэш-таблица, и порядок элементов зависит от хэш-функции и наличия коллизий.

## 8. В чем разница make и new?
- ```make``` используется для создания слайсов, мап, и каналов, которые являются ссылочными типами данных в Go. Например, чтобы создать слайс ```mySlice := make([]int, 0)```. Функция ```make``` инициализирует и возвращает указатель на созданный объект, и предоставляет специальные возможности для настройки внутренней реализации ссылочных типов.
- ```new``` используется для выделения памяти для нового значения указанного типа и возвращает указатель на выделенную память. Однако сама память остается неинициализированной, и требуется явно проинициализировать переменную, на которую указывает возвращенный указатель. Например: ```a := new(int)``` выделяет память для int и возвращает указатель на него.

## 9. Сколько существует способов задать переменную типа slice или map?
1. С помощью литералов:

- Для slice: 
```go 
    s := []int{1, 2, 3} 
```
- Для map:  
```go 
    m := map[string]int{"one": 1, "two": 2, "three": 3}
```

2. С использованием функций создания:

- Для slice:  
```go 
    s := make([]int, 0)
```

- Для map:  
```go 
    m := make(map[string]int)
```

3. С помощью оператора new:

- Для slice: 
```go 
    s := new([]int)
```

- Для map: 
```go 
    m := new(map[string]int)
```

4. С использованием инициализации значениями по умолчанию:

- Для slice: 
```go
    var s []int
```

- Для map: 
```go
    var m map[string]int
```

## 10. Что выведет данная программа и почему?

```go
    func update(p *int) {
        b := 2
        p = &b
    }

    func main() {
        var (
            a = 1
            p = &a
        )
        fmt.Println(*p)
        update(p)
        fmt.Println(*p)
    }
```
```
Out:
1
1
```
- В функции update, переменной ```p``` присваивается адрес переменной ```b```. Однако, это никак не влияет на исходную переменную ```p``` в функции ```main```. Поэтому, после вызова функции ```update(p)```, значение переменной ```p``` остается равным адресу переменной ```a```, а не адресу переменной ```b```. Поэтому вывод значений указателя ```p``` приведет к выводу значения переменной ```a```, которое равно 1, в обоих случаях.

## 11. Что выведет данная программа и почему?
```go
    func main() {
        wg := sync.WaitGroup{}
        for i := 0; i < 5; i++ {
            wg.Add(1)
            go func(wg sync.WaitGroup, i int) {
                fmt.Println(i)
                wg.Done()
            }(wg, i)
        }
        wg.Wait()
        fmt.Println("exit")
    }

```
```
Out:
4
0
1
2
3
fatal error: all goroutines are asleep - deadlock!
```
- Сначала числа от 0 до 4 в некотором порядке, а после ошибка deadlock.
- Ошибка находится там, где мы передаем переменную wg в анонимную функцию, мы фактически передаем копию объекта sync.WaitGroup. Из-за этого, каждая горутина работает с собственной копией wg, и вызов wg.Done() фактически не сокращает ожидание главной горутины. Внешняя структура никогда не получит сигнал о завершении 5 созданных горутин. 
- Можно исправить программу, передавая указатель на sync.WaitGroup вместо копии:
```go
    go func(wg *sync.WaitGroup, i int) {
        fmt.Println(i)
        wg.Done()
    }(&wg, i)
```

## 12. Что выведет данная программа и почему?
```go
    func main() {
        n := 0
        if true {
            n := 1
            n++
        }
        fmt.Println(n)
    }
```
```
Out:
0
```
- Объявление ```n := 1``` внутри блока if создает новую переменную ```n```, локальную для данного блока. Изменения, внесенные в эту локальную переменную ```n```, не будут влиять на глобальную переменную, объявленную в начале функции ```main```.

## 13. Что выведет данная программа и почему?
```go
    func someAction(v []int8, b int8) {
        v[0] = 100
        v = append(v, b)
    }

    func main() {
        var a = []int8{1, 2, 3, 4, 5}
        someAction(a, 6)
        fmt.Println(a)
    }
```
```
Out:
[100 2 3 4 5]
```
- Программа изменяет элемент с индексом 0 в срезе ```a``` на значение 100 внутри функции ```someAction```, поскольку срезы это ссылочный тип данных. Однако, при выполнении операции ```append```, создается новый срез с добавленным элементом ```b```, и изменения этого нового среза не отражаются на исходном срезе ```a```. Поэтому, изменение значения индекса 0 сохраняется, но добавление элемента 6 не отражается в исходном срезе ```a```.

## 14. Что выведет данная программа и почему?
```go
    func main() {
        slice := []string{"a", "a"}

        func(slice []string) {
            slice = append(slice, "a")
            slice[0] = "b"
            slice[1] = "b"
            fmt.Print(slice)
        }(slice)
        fmt.Print(slice)
    }
```
```
Out:
[b b a][a a]
```
- При вызове функции внутри функции (func(slice []string)) создается новая область видимости для slice, которая является копией оригинального среза.
- Внутри этой функции:
    1. append(slice, "a") добавляет новый элемент "a" к копии среза slice, но не изменяет оригинальный срез. Таким образом, копия среза становится ["a", "a", "a"].
    2. slice[0] = "b" изменяет первый элемент копии среза на "b". Копия среза теперь равна ["b", "a", "a"].
    3. slice[1] = "b" изменяет второй элемент копии среза на "b". Копия среза становится ["b", "b", "a"].
    4. fmt.Print(slice) выводит копию среза ["b", "b", "a"].
- Затем, возвращаясь к основной функции, она продолжает выполнение с оригинальным срезом slice, который не был изменен внутри вложенной функции. Поэтому fmt.Print(slice) выводит оригинальный срез ["a", "a"].
