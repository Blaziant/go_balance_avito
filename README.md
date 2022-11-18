# Golang Balance Avito

## Local start

- <code>docker-compose up --build</code>

## Описание методов пользователя:
- <code>/user/</code> - создать пользователя
  ```json
  { 
    "balance" : 100000.0,
  } 
  ```
- <code>/user/:id/balance</code> - возвращает баланс
- <code>/user/debit_balance</code> -  отнять сумму от баланса
  ```json
  { 
    "amount" : 10.0,
    "user_id" : 1
  } 
  ```
- <code>/user/replenish_balance</code> - добавить сумму к балансу
  ```json
  { 
    "amount" : 11.1,
    "user_id" : 1
  } 
  ```
- <code>/user/transfer_money</code> - перевод баланса от одного пользователя к другому
  ```json
  { 
    "from_id" : 1,
    "to_id" : 2,
    "amount" : 100.0
  }
  ```

## Описание методов услуги:
- <code>/favour/</code> - создать услугу
  ```json
  { 
    "name" : "колонка",
    "price" : 10000.0
  } 
  ```
- <code>/favour/:id</code> - получить услугу по id

## Описание методов заказа:
- <code>/order/</code> - создать заказ
  ```json
  { 
    "favours_id" : [1,2,3],
    "user_id": 1
  } 
  ```
- <code>/order/reserve_money</code> - резервация средств
  ```json
  { 
    "order_id" : 1 
  } 
  ```
- <code>/order/accept</code> - принятие выручки 
  ```json
  { 
    "order_id" : 1 
  } 
  ```
- <code>/order/report</code> - отчет для бухгалтерии
  ```json
  {
    "begin" : "2021-11-18", 
    "end" : "2023-11-18"
  } 
  ```