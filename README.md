# lms_calculator
**lms_calculator** - веб-сервис для вычисления результата выражения.

---

## Установка

1. установите go на компьютер, если он отсутствует
2. склонируйте репозиторий
```sh
git clone https://github.com/ThisIsHyum/lms_calculator
```
3. войдите в папку репозитория
```sh
cd lms_calculator
```
4. установите docker на компьютер, если он отсутствует
5. соберите проект
- **Linux**
```sh
sudo docker-compose build
```
- **Windows**
```sh
docker-compose build
```

## Запуск
1. войдите в папку репозитория
```sh
cd lms_calculator
```
2. запустите контейнеры(должен быть установлен docker)
- **Linux**
```sh
sudo docker-compose up
```
- **Windows**
```sh
docker-compose up
```
### параметры

Для того, чтобы изменить настройки, измените файл config.toml:
```toml
ip="localhost" # айпи
port=80 # порт

time_addition_ms=1 # длительность выполнения сложения в миллисекундах
time_subtraction_ms=1 # длительность выполнения вычитания в миллисекундах
time_multiplications_ms=1 # длительность выполнения умножения в миллисекундах
time_divisions_ms=1 # длительность выполнения деления в миллисекундах

ComputingPower=1 # количество горутин(агент)
```

## Запуск и установка

```sh
git clone https://github.com/ThisIsHyum/lms_calculator
cd lms_calculator
docker-compose build
docker-compose up
```

## Использование

1. создание выражения
- **url-путь**: `POST`
- **метод**: `/api/v1/calculate`
- **заголовок**: `Content-Type: application/json`
- **тело запроса**:
```json
{
  "expression": "выражение"
}
```
- **тело успешного ответа**:
```created```

### Примеры
1. **Выражение принято для вычисления**
    - **Статус**: `201`
    - **Пример запроса**
    ```sh
    curl --location 'http://localhost:8080/api/v1/calculate' \
    --header 'Content-Type: application/json' \
    --data '{
      "expression": "2+2*2"
    }'
    ```
    - **Пример ответа**
    ```created```
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
    ```error: brace error```

---
2. получение списка выражений

- **url-путь**: `/api/v1/expression`
- **метод**: `GET`
- **тело успешного ответа**:
```json
{
  "expressions": [
    {
      "id": "айди выражения",
      "status": "статус вычисления выражения",
      "result": "результат выражения"
    },
    {
      "id": "айди выражения",
      "status": "статус вычисления выражения",
      "result": "результат выражения"
    }
  ]
}
```

### Примеры
1. **Успешный запрос**
    - **Статус**: `200`
    - **Пример запроса**
    ```sh
    curl --location 'http://localhost:8080/api/v1/expressions'
    ```
    - **Пример ответа**
    ```json
    {
      "expressions": [
        {
          "id": 1,
          "status": "solved",
          "result": 4
        },
        {
          "id": 2,
          "status": "not resolved",
          "result": 0
        }
      ]
    }
    ```
---
3. получение выражения

- **url-путь**: `/api/v1/expressions/:id`
- **метод**: `GET`

- **тело успешного ответа**:
```json
{
  "expression":
    {
      "id": "айди выражения",
      "status": "статус вычисления выражения",
      "result": "результат выражения"
    }
}
```

### Примеры
1. **Успешный запрос**
    - **Статус**: `200`
    - **Пример запроса**
     ```sh
    curl --location 'http://localhost:8080/api/v1/expressions/1'
    ```
    - **Пример ответа**
    ```json
    {
      "expression":
        {
          "id": 1,
          "status": "solved",
          "result": 4
        }
    }
    ```
2. **Нет задачи**
    - **Статус**: `404`
    - **Пример запроса**
     ```sh
    curl --location 'http://localhost:8080/api/v1/expressions/3'
    ```
    - **Пример ответа**
    ```Not found```

## Схема работы
```
|-------|  POST/GET  |--------------|              POST/GET                 |-------|
| user  | ---------> | orchestrator | <------------------------------------ | agent |
|-------|            |--------------|  [computingPower] requests in 1 second  |-------|
```
