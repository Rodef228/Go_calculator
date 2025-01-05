# Go Calculator Service

## Привет, это Rodef228

### Описание
HTTP_Calculator — это простой HTTP-сервис для вычисления математических выражений. Сервис принимает POST-запросы с математическими выражениями и возвращает результат вычислений в формате JSON.

### Запуск
```bash
go run ./cmd/calc_service/...
```
Не забудьте, надо находиться в корневой папке!

### Использование API с помощью curl

#### Эндпоинт
URL: http://localhost:8080/api/v1/calculate<br>
Метод: POST<br>
Заголовок: Content-Type: application/json<br>
Тело запроса:<br>
```json
{
    "expression": "expression"
}
```


#### Примеры запросов (Windows)
1. Успешный подсчёт (200)
```bash
curl --location 'http://localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '{"expression":"2+2*2"}'
```

Результат:
```json
{
    "result": 6
}
```


2. Некорректное выражение (422)
```bash
curl -X POST http://localhost:8080/api/v1/calculate ^ -H "Content-Type: application/json" ^ -d "{\"expression\": \"2+a\"}"
```

Результат:
```json
{
    "error": "Invalid expression format"
}
```


3. Неправильный формат JSON (500)
```bash
curl -X POST http://localhost:8080/api/v1/calculate ^ -H "Content-Type: application/json" ^ -d "invalid json"
```

Результат:
```json
{
    "error": "Internal server error"
}
```

### Запуск тестов
В проекте уже реализованы автоматические тесты для калькулятора и HTTP-обработчиков.
1. Запустите тесты калькулятора:
```bash
go test ./internal/calculator
```
2. Запустите тесты обработчиков:
```bash
go test ./internal/handlers
```
3. Запустите все тесты одновременно:
```bash
go test ./...
```

### Вопросы?
Для вопросов и предложений, пожалуйста, создайте issue на GitHub или напишите в телегу: Rodef228
