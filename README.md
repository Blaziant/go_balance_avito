# Golang Balance Avito

## Local start

- <code>docker-compose up --build</code>

## Пользователь:
- <code>/user/</code> - создать пользователя
- <code>/user/:id/balance</code> - возвращает баланс 
- <code>/user/debit_balance</code> -  отнять сумму от баланса
- <code>/user/replenish_balance</code> - добавить сумму к балансу
- <code>/user/transfer_money</code> - перевод баланса от одного пользователя к другому

## Услуга
Описание методов:
- <code>/favour/</code> - создать услугу
- <code>/favour/:id</code> - получить услугу по id

## Заказ
Описание методов:
- <code>/order/</code> - создать заказ
- <code>/order/reserve_money</code> - резервация средств
- <code>/order/accept</code> - принятие заказа, принятие выручки