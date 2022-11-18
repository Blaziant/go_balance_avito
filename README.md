# Golang Balance Avito

## Local start

- <code>docker-compose up --build</code>

## Описание методов пользователя:
- GET <code>/api/v1/user/:id/balance</code> - возвращает баланс
- POST <code>/api/v1/user/</code> - создать пользователя
  ```json
  { 
    "balance" : "Баланс пользователя float64",
  } 
  ```
- POST <code>/api/v1/user/debit_balance</code> -  отнять сумму от баланса
  ```json
  { 
    "amount" : "На сколько нужно уменьшить float64",
    "user_id" : "id пользователя uint"
  } 
  ```
- POST <code>/api/v1/user/replenish_balance</code> - добавить сумму к балансу
  ```json
  { 
    "amount" : "На сколько нужно увеличить float64",
    "user_id" : "id пользователя uint"
  } 
  ```
- POST <code>/api/v1/user/transfer_money</code> - перевод баланса от одного пользователя к другому
  ```json
  { 
    "from_id" : "id пользователя откуда переводить uint",
    "to_id" : "id пользователя куда переводить uint",
    "amount" : "Сколько нужно перевести средств float64"
  }
  ```

## Описание методов услуги:
- GET <code>/api/v1/favour/:id</code> - получить услугу по id
- POST <code>/api/v1/favour/</code> - создать услугу
  ```json
  { 
    "name" : "Название услуги string",
    "price" : "Цена услуги float64"
  } 
  ```

## Описание методов заказа:
- POST <code>/api/v1/order/</code> - создать заказ
  ```json
  { 
    "favours_id" : "Массив индексов услуг в этом заказе []uint",
    "user_id": "id пользователя, который сделал заказ uint"
  } 
  ```
- POST <code>/api/v1/order/reserve_money</code> - резервация средств
  ```json
  { 
    "order_id" : "id заказа для резервирования средств uint"
  } 
  ```
- POST <code>/api/v1/order/accept</code> - принятие выручки 
  ```json
  { 
    "order_id" : "id заказа для принятия заказа uint"
  } 
  ```
- GET <code>/api/v1/order/report</code> - отчет для бухгалтерии
  ```json
  {
    "begin" : "Дата начала для отчёта в виде (2021-11-18) string", 
    "end" : "Дата начала для отчёта в виде (2023-11-18) string"
  } 
  ```