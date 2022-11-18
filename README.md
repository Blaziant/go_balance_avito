# Golang Balance Avito

## Local start

- <code>docker-compose up --build</code>

## Описание методов пользователя
/user/ - создать пользователя
/user/:id/balance - возвращает баланс 
/user/debit_balance -  отнять сумму от баланса
/user/replenish_balance - добавить сумму к балансу
/user/transfer_money - перевод баланса от одного пользователя к другому

## Описание методов услуги
/favour/ - создать услугу
/favour/:id - получить услугу по id

## Описание методов заказа
/order/ - создать заказ
/order/reserve_money - резервация средств
/order/accept - принятие заказа, принятие выручки