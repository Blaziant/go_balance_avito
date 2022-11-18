# Golang Balance Avito

## Local start

- <code>docker-compose up --build</code>

## Описание методов пользователя:
- <code>/api/v1/user/:id/balance</code> - возвращает баланс
- <code>/api/v1/user/</code> - создать пользователя
  ```json
  { 
    "balance" : 100000.0,
  } 
  ```
- <code>/api/v1/user/debit_balance</code> -  отнять сумму от баланса
  ```json
  { 
    "amount" : 10.0,
    "user_id" : 1
  } 
  ```
- <code>/api/v1/user/replenish_balance</code> - добавить сумму к балансу
  ```json
  { 
    "amount" : 11.1,
    "user_id" : 1
  } 
  ```
- <code>/api/v1/user/transfer_money</code> - перевод баланса от одного пользователя к другому
  ```json
  { 
    "from_id" : 1,
    "to_id" : 2,
    "amount" : 100.0
  }
  ```

## Описание методов услуги:
- <code>/api/v1/favour/:id</code> - получить услугу по id
- <code>/api/v1/favour/</code> - создать услугу
  ```json
  { 
    "name" : "колонка",
    "price" : 10000.0
  } 
  ```

## Описание методов заказа:
- <code>/api/v1/order/</code> - создать заказ
  ```json
  { 
    "favours_id" : [1,2,3],
    "user_id": 1
  } 
  ```
- <code>/api/v1/order/reserve_money</code> - резервация средств
  ```json
  { 
    "order_id" : 1 
  } 
  ```
- <code>/api/v1/order/accept</code> - принятие выручки 
  ```json
  { 
    "order_id" : 1 
  } 
  ```
- <code>/api/v1/order/report</code> - отчет для бухгалтерии
  ```json
  {
    "begin" : "2021-11-18", 
    "end" : "2023-11-18"
  } 
  ```