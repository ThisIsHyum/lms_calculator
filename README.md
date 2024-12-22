# lms_calculator
**lms_calculator** - веб-сервис для вычисления результата выражения.

--

## Установка

1. установите go на компьютер, если он отсутствует
2. склонируйте репозиторий
```sh
git clone https://github.com/ThisIsHyum/lms_calculator
```

## Запуск

1. войдите в папку репозитория
```sh
cd lms_calculator
```
2. запустите проект(с root-правами, если используется порт со значением до 1024)
```sh
sudo go run cmd/calc_service/.
```

### параметры

Для того, чтобы присвоить айпи и порту свои значения, необходимо использовать флаги ip и port:
```sh
go run cmd/calc_service/. -ip 192.168.0.100 -port 8080
```

## Запуск и установка

```sh
git clone https://github.com/ThisIsHyum/lms_calculator
cd lms_calculator
sudo go run cmd/calc_service/.
```

## Использование

- **метод**: `POST`
  
- url-путь: `/api/v1/calculate`
  
- заголовок: `Content-Type: application/json`
  
- тело запроса:
```json
{
  "expression": "выражение"
}
```

### Ответы
1. **Успешный запрос**
    - **Статус**: `200`
    - **Пример запроса**
    ```sh
    curl --location 'http://localhost:8080/api/v1/calculate' \
    --header 'Content-Type: application/json' \
    --data '{
      "expression": "2+2*2"
    }'
    ```
    - **Пример ответа**
    ```json
    {
      "result": "6"
    }
    ```
2. **Некорректное выражение**
    - **Статус**: `422`
    - **Пример запроса**
    ```sh
    curl --location 'http://localhost:8080/api/v1/calculate' \
    --header 'Content-Type: application/json' \
    --data '{
      "expression": "1+(2"
    }'
    ```
    - **Пример ответа**
    ```json
    {
      "error": "bracket not closed"
    }
    ```
3. **Некорректный метод**
    - **Статус**: `405`
    - **Пример запроса**
    ```sh
    curl --location 'http://localhost:8080/api/v1/calculate'
    ```
    - **Пример ответа**
    ```json
    {
      "error": "Wrong method"
    }
    ```
