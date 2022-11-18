# Golang Balance Avito

## Local start

- <code>docker-compose up --build</code>

## Описание методов пользователя:
- GET <code>/api/v1/user/:id/balance</code> - возвращает баланс
- <code>/api/v1/user/</code> - создать пользователя
  ```json
  { 
    "balance" : "Баланс пользователя",
  } 
  ```
- <code>/api/v1/user/debit_balance</code> -  отнять сумму от баланса
  ```json
  { 
    "amount" : "На сколько нужно уменьшить",
    "user_id" : "id пользователя"
  } 
  ```
- <code>/api/v1/user/replenish_balance</code> - добавить сумму к балансу
  ```json
  { 
    "amount" : "На сколько нужно увеличить",
    "user_id" : "id пользователя"
  } 
  ```
- <code>/api/v1/user/transfer_money</code> - перевод баланса от одного пользователя к другому
  ```json
  { 
    "from_id" : "id пользователя откуда переводить",
    "to_id" : "id пользователя куда переводить",
    "amount" : "Сколько нужно перевести средств"
  }
  ```

## Описание методов услуги:
- <code>/api/v1/favour/:id</code> - получить услугу по id
- <code>/api/v1/favour/</code> - создать услугу
  ```json
  { 
    "name" : "Название услуги",
    "price" : "Цена услуги"
  } 
  ```

## Описание методов заказа:
- <code>/api/v1/order/</code> - создать заказ
  ```json
  { 
    "favours_id" : "Массив индексов услуг в этом заказе",
    "user_id": "id пользователя, который сделал заказ"
  } 
  ```
- <code>/api/v1/order/reserve_money</code> - резервация средств
  ```json
  { 
    "order_id" : "id заказа для резервирования средств"
  } 
  ```
- <code>/api/v1/order/accept</code> - принятие выручки 
  ```json
  { 
    "order_id" : "id заказа для принятия заказа"
  } 
  ```
- <code>/api/v1/order/report</code> - отчет для бухгалтерии
  ```json
  {
    "begin" : "Дата начала для отчёта в виде строки (2021-11-18)", 
    "end" : "Дата начала для отчёта в виде строки (2023-11-18)"
  } 
  ```