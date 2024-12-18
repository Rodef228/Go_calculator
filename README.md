Привет, я Rodef228.

Что делает серв, он получает выражение, и возвращает рещультат.

Чтобы запустить сервер надо прописать `go run main.go`.
Теперь протестировать надо:
url: **http://localhost:8080/api/v1/calculate**
headers:  **Content-Type: application/json**
body: 
{
  "expression": "52+52"
}

результатом приходит:
{
    "result": "104"
}

если введено неверно, то status code = 422
если ошибка в символах или /0, то 500
если всё хорошо, то 200

примеры ввода и возможные исходы:

1. {"expression":"52+52"}
    status code 200
    {"result":"104"}

2. {"expression": "52/0"}
    status code 500
    {"error":"Internal server error"}

3. {"expression": ""}
    status code 422
    {"error":"Expression is not valid"}

как-то так :)