# Avito test task
Попытка выполнения тестового задания
# Инструкция по запуску
```bash
docker-compose up --build
```
# API
Данный микросервис имеет четыре метода. Примеры запросов можно найти в файле avito-test.postman_collection.json
## Метод получения баланса пользователя
### Post "/user/"
Данный метод принимает id пользователя и возвращает баланс пользователя.
### Пример
```bash
curl --location --request GET 'http://localhost:8800/user/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "userId": 10
}'
```
## Метод начисления
### Get "/user/"
Данный метод принимает id пользователя и сумму необходиммую для начисления на баланс.
### Пример
```bash
curl --location --request POST 'http://localhost:8800/user/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "userId": 3,
    "balance": 400
}'
```
## Метод резервирования средств
### Post "/order/reserve"
Данный метод создаёт заказ, резервируя средства пользователя на отдельном счете.
### Пример
```bash
curl --location --request POST '127.0.0.1:8080/order/reserve' \
--header 'Content-Type: application/json' \
--data-raw '{
    "orderId": 0,
    "serviceId": 14125125,
    "userId": 10,
    "price": 1000
}'
```
## Метод признания выручки
Данный метод переводит заказ в статус выполненного, списывая зарезервированные средства пользователя.
### Post "/order/revenue"
### Пример
```bash
curl --location --request POST 'http://localhost:8800/order/revenue' \
--header 'Content-Type: application/json' \
--data-raw '{
    "orderId": 0,
    "serviceId": 14125125,
    "userId": 10,
    "price": 1000
}'
```

