# Cast - Type Conversion Library for Go

Библиотека `cast` предоставляет удобные функции для преобразования между различными типами данных в Go. Все функции поддерживают обработку `nil` значений и предлагают два варианта использования: с паникой при ошибке (функции `To*`) и с возвращением ошибки (функции `As*`).

## Содержание

- [Базовые типы](#базовые-типы)
  - [Boolean](#boolean)
  - [Integer](#integer)
  - [Unsigned Integer](#unsigned-integer)
  - [Floating Point](#floating-point)
  - [String](#string)
- [Сложные типы](#сложные-типы)
  - [Slice](#slice)
  - [Map](#map)
  - [Pointer](#pointer)
  - [Interface](#interface)
  - [Struct](#struct)
- [Специальные типы](#специальные-типы)
  - [Time](#time)
  - [UUID](#uuid)
- [Универсальные функции](#универсальные-функции)
  - [AsType / ToType](#astype--totype)
  - [Converter](#converter)

---

## Базовые типы

### Boolean

#### `ToBool(value any) bool`
Преобразует значение в `bool`. При ошибке вызывает panic.

**Аргументы:**
- `value any` - значение для преобразования

**Возвращает:**
- `bool` - преобразованное значение

**Поддерживаемые типы:**
- `bool` - возвращается как есть
- `int`, `uint`, `int8-64`, `uint8-64` - не равно 0 = true
- `float32`, `float64` - не равно 0 = true
- `string` - "true" = true, "false" или пустая строка = false
- указатели и интерфейсы - распаковываются и преобразуются рекурсивно
- `nil` - возвращает false

**Пример:**
```go
result := cast.ToBool(1)           // true
result := cast.ToBool(0)           // false
result := cast.ToBool("true")      // true
result := cast.ToBool("false")     // false
result := cast.ToBool(nil)         // false
```

#### `AsBool(value any) (bool, error)`
Преобразует значение в `bool`. Возвращает ошибку при невозможности преобразвания.

**Аргументы:**
- `value any` - значение для преобразования

**Возвращает:**
- `bool` - преобразованное значение
- `error` - ошибка при преобразовании или `nil`

**Пример:**
```go
result, err := cast.AsBool("invalid")
if err != nil {
    log.Fatal(err)  // failed to cast "invalid" to bool
}
```

---

### Integer

#### `ToInt(value any) int`
Преобразует значение в `int`. При ошибке вызывает panic.

**Аргументы:**
- `value any` - значение для преобразования

**Возвращает:**
- `int` - преобразованное значение

**Поддерживаемые типы:**
- `int`, `uint`, `int8-64`, `uint8-64` - прямое преобразование
- `float32`, `float64` - округление вниз (truncation)
- `bool` - true = 1, false = 0
- `string` - парсирование через `strconv.Atoi()`
- `nil` - возвращает 0

**Пример:**
```go
result := cast.ToInt(42)           // 42
result := cast.ToInt("123")        // 123
result := cast.ToInt(123.7)        // 123
result := cast.ToInt(true)         // 1
result := cast.ToInt(nil)          // 0
```

#### `AsInt(value any) (int, error)`
Преобразует значение в `int`. Возвращает ошибку при невозможности преобразования.

**Аргументы:**
- `value any` - значение для преобразования

**Возвращает:**
- `int` - преобразованное значение
- `error` - ошибка при преобразовании или `nil`

---

### Int8, Int16, Int32, Int64

Аналогичны функциям для `int`, но для различных размеров целых чисел со знаком.

#### `ToInt8(value any) int8` / `AsInt8(value any) (int8, error)`
#### `ToInt16(value any) int16` / `AsInt16(value any) (int16, error)`
#### `ToInt32(value any) int32` / `AsInt32(value any) (int32, error)`
#### `ToInt64(value any) int64` / `AsInt64(value any) (int64, error)`

**Пример:**
```go
result := cast.ToInt8("100")       // 100
result := cast.ToInt16(32000)      // 32000
result := cast.ToInt32("1000000")  // 1000000
result := cast.ToInt64(true)       // 1
```

---

### Unsigned Integer

#### `ToUint(value any) uint`
Преобразует значение в `uint`. При ошибке вызывает panic.

**Аргументы:**
- `value any` - значение для преобразования

**Возвращает:**
- `uint` - преобразованное значение

**Поддерживаемые типы:** аналогичны `ToInt`, но для беззнаковых чисел.

**Пример:**
```go
result := cast.ToUint(42)          // 42
result := cast.ToUint("123")       // 123
result := cast.ToUint(123.7)       // 123
result := cast.ToUint(nil)         // 0
```

### Uint8, Uint16, Uint32, Uint64

Аналогичны функциям для `uint`, но для различных размеров беззнаковых чисел.

#### `ToUint8(value any) uint8` / `AsUint8(value any) (uint8, error)`
#### `ToUint16(value any) uint16` / `AsUint16(value any) (uint16, error)`
#### `ToUint32(value any) uint32` / `AsUint32(value any) (uint32, error)`
#### `ToUint64(value any) uint64` / `AsUint64(value any) (uint64, error)`

**Пример:**
```go
result := cast.ToUint8("200")      // 200
result := cast.ToUint16(60000)     // 60000
result := cast.ToUint32("4000000") // 4000000
result := cast.ToUint64(false)     // 0
```

---

### Floating Point

#### `ToFloat32(value any) float32`
Преобразует значение в `float32`. При ошибке вызывает panic.

**Аргументы:**
- `value any` - значение для преобразования

**Возвращает:**
- `float32` - преобразованное значение

**Поддерживаемые типы:**
- целые числа - преобразование в float
- `float32`, `float64` - прямое преобразование
- `bool` - true = 1.0, false = 0.0
- `string` - парсирование через `strconv.ParseFloat()`
- `nil` - возвращает 0.0

**Пример:**
```go
result := cast.ToFloat32(42)       // 42.0
result := cast.ToFloat32("123.45") // 123.45
result := cast.ToFloat32(true)     // 1.0
result := cast.ToFloat32(nil)      // 0.0
```

#### `ToFloat64(value any) float64` / `AsFloat64(value any) (float64, error)`
Аналого `ToFloat32` / `AsFloat32`, но для `float64`.

**Пример:**
```go
result := cast.ToFloat64(42)       // 42.0
result := cast.ToFloat64("123.456789") // 123.456789
result, err := cast.AsFloat64("3.14")
```

---

### String

#### `ToString(value any) string`
Преобразует значение в строку. При ошибке вызывает panic.

**Аргументы:**
- `value any` - значение для преобразования

**Возвращает:**
- `string` - преобразованное значение

**Поддерживаемые типы:**
- `string` - возвращается как есть
- `bool` - "true" или "false"
- `int`, `uint`, целые числа - через `strconv` функции
- `float32`, `float64` - через `strconv.FormatFloat()`
- типы с методом `String()` (реализующие `fmt.Stringer`)
- указатели и интерфейсы - распаковываются и преобразуются
- `nil` - возвращает пустую строку

**Пример:**
```go
result := cast.ToString(42)        // "42"
result := cast.ToString(3.14)      // "3.14"
result := cast.ToString(true)      // "true"
result := cast.ToString(nil)       // ""
result := cast.ToString(time.Now()) // строка с временем
```

#### `AsString(value any) (string, error)`
Преобразует значение в строку. Возвращает ошибку при невозможности преобразования.

**Пример:**
```go
result, err := cast.AsString(42)
```

---

## Сложные типы

### Slice

#### `ToSlice[V any](value any) []V`
Преобразует значение в срез типа `[]V`. При ошибке вызывает panic.

**Аргументы:**
- `value any` - значение для преобразования

**Возвращает:**
- `[]V` - преобразованный срез

**Поддерживаемые типы:**
- срезы - каждый элемент преобразуется в тип `V`
- указатели на срезы
- `nil` - возвращает `nil`

**Пример:**
```go
// Преобразование срезов разных типов
result := cast.ToSlice[int]([]string{"1", "2", "3"})
// result: []int{1, 2, 3}

result := cast.ToSlice[string]([]int{1, 2, 3})
// result: []string{"1", "2", "3"}

result := cast.ToSlice[int64]([]any{1, "2", 3.0})
// result: []int64{1, 2, 3}
```

#### `AsSlice[V any](value any) ([]V, error)`
Преобразует значение в срез типа `[]V`. Возвращает ошибку если преобразование невозможно.

**Пример:**
```go
result, err := cast.AsSlice[int]([]string{"1", "2", "3"})
if err != nil {
    log.Fatal(err)
}
```

---

### Map

#### `ToMap[K comparable, V any](value any) map[K]V`
Преобразует значение в маппу типа `map[K]V`. При ошибке вызывает panic.

**Аргументы:**
- `value any` - значение для преобразования

**Возвращает:**
- `map[K]V` - преобразованная маппа

**Поддерживаемые типы:**
- маппы - ключи и значения преобразуются в типы `K` и `V`
- указатели на маппы
- `nil` - возвращает `nil`

**Пример:**
```go
// Преобразование маппы со строковыми ключами и значениями
originalMap := map[string]string{
    "count": "10",
    "price": "99.99",
}
result := cast.ToMap[string, int](originalMap)
// result: map[string]int{"count": 10, "price": 99}

// Преобразование маппы с разными типами ключей
mixedMap := map[any]any{
    1:     "one",
    "two": 2,
}
result := cast.ToMap[string, string](mixedMap)
```

#### `AsMap[K comparable, V any](value any) (map[K]V, error)`
Преобразует значение в маппу типа `map[K]V`. Возвращает ошибку если преобразование невозможно.

**Пример:**
```go
result, err := cast.AsMap[string, int](map[string]string{"a": "1", "b": "2"})
if err != nil {
    log.Fatal(err)
}
```

---

### Pointer

#### `ToPointer[V any](value any) *V`
Преобразует значение в указатель на `V`. При ошибке вызывает panic.

**Аргументы:**
- `value any` - значение для преобразования

**Возвращает:**
- `*V` - указатель на преобразованное значение

**Поддерживаемые типы:**
- уже готовые указатели типа `*V`
- любые значения, которые могут быть преобразованы в `V`
- `nil` - возвращает `nil`

**Пример:**
```go
// Создание указателя на int из строки
result := cast.ToPointer[int]("42")
// result: &42 (указатель на int)

// Создание указателя на строку из числа
result := cast.ToPointer[string](123)
// result: &"123" (указатель на строку)

// nil остается nil
result := cast.ToPointer[int](nil)
// result: nil
```

#### `AsPointer[V any](value any) (*V, error)`
Преобразует значение в указатель на `V`. Возвращает ошибку если преобразование невозможно.

**Пример:**
```go
result, err := cast.AsPointer[float64]("3.14")
if err != nil {
    log.Fatal(err)
}
```

---

### Interface

#### `ToInterface[V any](value any) V`
Преобразует значение в интерфейс типа `V`. При ошибке вызывает panic.

**Аргументы:**
- `value any` - значение для преобразования

**Возвращает:**
- `V` - преобразованное значение (интерфейс)

**Пример:**
```go
// Определение интерфейса
type Reader interface {
    Read(p []byte) (n int, err error)
}

// Преобразование значения в интерфейс
value := strings.NewReader("Hello")
result := cast.ToInterface[Reader](value)
```

#### `AsInterface[V any](value any) (V, error)`
Преобразует значение в интерфейс типа `V`. Возвращает ошибку если преобразование невозможно.

---

### Struct

#### `ToStruct[S any](value any) S`
Преобразует значение в структуру типа `S`. При ошибке вызывает panic.

**Аргументы:**
- `value any` - значение для преобразования (маппа или другая структура)

**Возвращает:**
- `S` - преобразованная структура

**Примечания:**
- Преобразует маппы в структуры, сопоставляя ключи маппы с полями структуры
- Поддерживает структуры с разными вариантами имен полей: точное совпадение, с заглавной или строчной буквой
- Преобразует значения полей если требуется
- `nil` - возвращает нулевую структуру

**Пример:**
```go
type Person struct {
    Name string
    Age  int
}

// Из маппы
data := map[string]any{
    "name": "John",
    "age":  30,
}
person := cast.ToStruct[Person](data)
// Person{Name: "John", Age: 30}

// Из другой структуры
type PersonDTO struct {
    Name string
    Age  string // строка вместо int
}
dto := PersonDTO{Name: "Alice", Age: "25"}
person := cast.ToStruct[Person](dto)
// Person{Name: "Alice", Age: 25}
```

#### `AsStruct[S any](value any) (S, error)`
Преобразует значение в структуру типа `S`. Возвращает ошибку если преобразование невозможно.

**Пример:**
```go
type Config struct {
    Host string
    Port int
}

data := map[string]any{"host": "localhost", "port": "8080"}
config, err := cast.AsStruct[Config](data)
if err != nil {
    log.Fatal(err)
}
```

---

## Специальные типы

### Time

#### `ToTime(value any) time.Time`
Преобразует значение в `time.Time`. При ошибке вызывает panic.

**Аргументы:**
- `value any` - значение для преобразования (обычно строка или `time.Time`)

**Возвращает:**
- `time.Time` - преобразованное время

**Поддерживаемые форматы:**
- `time.Kitchen` - "3:04PM"
- `time.DateOnly` - "2006-01-02"
- `time.TimeOnly` - "15:04:05"
- `time.DateTime` - "2006-01-02 15:04:05"
- `time.RFC3339` - "2006-01-02T15:04:05Z07:00"
- Кастомные форматы (см. `TimeLayouts`)
- Уже готовые `time.Time` - возвращаются как есть

**Пример:**
```go
result := cast.ToTime("2024-01-15")
// result: время 15 января 2024 года

result := cast.ToTime("2024-01-15T10:30:45Z")
// result: время 15 января 2024 года 10:30:45 UTC

result := cast.ToTime("3:04PM")
// result: время сегодня в 3:04 PM

result := cast.ToTime(time.Now())
// result: текущее время
```

#### `AsTime(value any) (time.Time, error)`
Преобразует значение в `time.Time`. Возвращает ошибку если преобразование невозможно.

**Пример:**
```go
result, err := cast.AsTime("2024-01-15")
if err != nil {
    log.Fatal(err) // failed to parse ... to time
}
```

#### `TimeLayouts`
Переменная, содержащая список поддерживаемых форматов времени для парсирования. Может использоваться или расширяться при необходимости.

**Пример:**
```go
// Просмотр поддерживаемых форматов
for _, layout := range cast.TimeLayouts {
    fmt.Println(layout)
}

// Расширение списка (если требуется)
customLayouts := append(cast.TimeLayouts, "02-01-2006")
```

---

### UUID

#### `ToUuid(value any) uuid.UUID`
Преобразует значение в `uuid.UUID`. При ошибке вызывает panic.

**Аргументы:**
- `value any` - значение для преобразования (строка, []byte или uuid.UUID)

**Возвращает:**
- `uuid.UUID` - преобразованный UUID

**Поддерживаемые типы:**
- `string` - парсирование строки UUID
- `[]byte` - парсирование байтового массива (16 байт или строка)
- `uuid.UUID` - возвращается как есть
- `nil` - возвращает `uuid.Nil`

**Пример:**
```go
// Из строки
result := cast.ToUuid("550e8400-e29b-41d4-a716-446655440000")
// result: uuid.UUID с значением

// Из байтового массива
bytes := []byte("550e8400-e29b-41d4-a716-446655440000")
result := cast.ToUuid(bytes)

// Из uuid.UUID
id := uuid.New()
result := cast.ToUuid(id)
// result: id

// nil
result := cast.ToUuid(nil)
// result: uuid.Nil (00000000-0000-0000-0000-000000000000)
```

#### `AsUuid(value any) (uuid.UUID, error)`
Преобразует значение в `uuid.UUID`. Возвращает ошибку если преобразование невозможно.

**Пример:**
```go
result, err := cast.AsUuid("invalid-uuid")
if err != nil {
    log.Fatal(err) // failed to cast ... to UUID
}
```

---

## Универсальные функции

### AsType / ToType

#### `AsType[V any](value any) (V, error)`
Универсальная функция для преобразования в любой тип. Возвращает ошибку при невозможности.

**Аргументы:**
- `value any` - значение для преобразования

**Возвращает:**
- `V` - преобразованное значение типа V
- `error` - ошибка при преобразовании или `nil`

**Пример:**
```go
// Преобразование в int
result, err := cast.AsType[int]("42")
if err != nil {
    log.Fatal(err)
}

// Преобразование в []string
result, err := cast.AsType[[]string]([]any{"a", "b", "c"})

// Преобразование в map[string]int
result, err := cast.AsType[map[string]int](map[string]string{"a": "1", "b": "2"})

// Преобразование в пользовательскую структуру
type Config struct {
    Host string
    Port int
}
result, err := cast.AsType[Config](map[string]any{"host": "localhost", "port": 8080})
```

#### `ToType[V any](value any) V`
Универсальная функция для преобразования в любой тип. При ошибке вызывает panic.

**Пример:**
```go
// Преобразование в int (паника при ошибке)
result := cast.ToType[int]("42")

// Преобразование в пользовательский тип
type User struct {
    ID   int
    Name string
}
user := cast.ToType[User](map[string]any{"id": 1, "name": "John"})
```

---

### Converter

#### `Converter[V any, C func(any) (V, error)]() (C, error)`
Возвращает функцию-конвертер для типа V, которая может быть использована многократно.

**Аргументы:** нет

**Возвращает:**
- `C` - функция конвертер типа `func(any) (V, error)`
- `error` - ошибка если конвертер для типа V не существует

**Примечание:**
- Это продвинутая функция для оптимизации при необходимости преобразовать много значений в один тип

**Пример:**
```go
// Получение конвертера для типа int
converter, err := cast.Converter[int]()
if err != nil {
    log.Fatal(err)
}

// Использование конвертера для преобразования множественных значений
values := []any{"1", "2", "3", "invalid"}
for _, val := range values {
    result, err := converter(val)
    if err != nil {
        fmt.Printf("Error converting %v: %v\n", val, err)
        continue
    }
    fmt.Printf("Converted: %d\n", result)
}
// Output:
// Converted: 1
// Converted: 2
// Converted: 3
// Error converting invalid: strconv.Atoi: parsing "invalid": invalid syntax

// Получение конвертера для сложного типа
sliceConverter, err := cast.Converter[[]int]()
if err != nil {
    log.Fatal(err)
}

// Использование
result, err := sliceConverter([]string{"1", "2", "3"})
if err != nil {
    log.Fatal(err)
}
fmt.Println(result) // [1 2 3]
```

---

## Рекомендации и советы

### Выбор между To* и As* функциями

- **To***: используйте когда вы уверены, что преобразование будет успешным, или хотите быстро обработать ошибку через panic
- **As***: используйте в production коде для явной обработки ошибок

```go
// Быстрый скрипт или тест
value := cast.ToInt("42")

// Production код
value, err := cast.AsInt(userInput)
if err != nil {
    return fmt.Errorf("invalid input: %w", err)
}
```

### Работа с nil значениями

Все функции безопасно работают с `nil`:
- Примитивные типы (`int`, `bool`, и т.д.) возвращают нулевое значение
- Сложные типы (`[]T`, `map[K]V`, `*T`) возвращают `nil`

```go
cast.ToInt(nil)        // 0
cast.ToString(nil)     // ""
cast.ToSlice[int](nil) // nil
cast.ToMap[string, int](nil) // nil
```

### Оптимизация при множественных преобразованиях

Для преобразования большого количества значений в один тип используйте `Converter`:

```go
// Неоптимально (функция вызывается для каждого значения)
for _, str := range strings {
    cast.AsInt(str)
}

// Оптимально (функция получена один раз)
converter, _ := cast.Converter[int]()
for _, str := range strings {
    converter(str)
}
```

### Обработка ошибок по типам

Разные типы могут выдавать разные ошибки:

```go
value, err := cast.AsInt("not-a-number")
// err: strconv.Atoi: parsing "not-a-number": invalid syntax

value, err := cast.AsStruct[Person](map[string]any{"age": "invalid"})
// err: failed to cast string to int for field conversion

value, err := cast.AsTime("invalid-date")
// err: failed to parse "invalid-date" to time
```

---

## Требования

- Go 1.18 или выше (использует Generics)
- `github.com/google/uuid` для работы с UUID
