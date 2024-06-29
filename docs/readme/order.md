# Order Service

- ### [1. Заказ](#1-заказ)
    - [Добавление заказа](#добавление-заказа)
    - [Получение всех заказов](#получение-всех-заказов)
    - [Удаление заказа](#удаление-заказа)
- ### [2. Предметы заказа](#2-предметы-заказа)
    - [Добавление предмета заказа](#добавление-предмета-заказа)
    - [Получение списка предметов заказа по order id](#получение-списка-предметов-заданного-заказа)
    - [Получение списка предметов заказа по product id](#получение-списка-предметов-заказов-заданного-продукта)
    - [Удаление предмета заказа](#удаление-заданного-предмета-заказа)

## 1. Заказ

### Добавление заказа

`POST` /order<br>

#### <span style="color:#12ff63">200 STATUS: OK

```json
{
  "id": 1
}
```

#### <span style="color:#df0000">400 STATUS: BAD REQUEST

```json
{
  "message": "invalid request"
}
```

### Получение всех заказов

`GET` /order<br>

#### <span style="color:#12ff63">200 STATUS: OK

```json
[
  {
    "id": 1,
    "datetime": "240629134303"
  },
  {
    "id": 2,
    "datetime": "240629134308"
  }
]
```

### Удаление заказа

`DELETE` /order/1<br>

#### <span style="color:#12ff63">200 STATUS: OK

```json
{
  "status": "ok"
}
```

#### <span style="color:#df0000">400 STATUS: BAD REQUEST

```json
{
  "message": "order with specified id not found"
}
```

## 2. Предметы заказа

### Добавление предмета заказа

`POST` /order/item<br>
Ожидаемый формат входных данных:

```json
{
  "order_id": 1,
  "product_id": 2,
  "quantity": 10
}
```

#### <span style="color:#12ff63">200 STATUS: OK

```json
[
  {
    "id": 1
  },
  {
    "id": 2
  }
]
```

#### <span style="color:#df0000">400 STATUS: BAD REQUEST

```json
{
  "message": "invalid request"
}
```

### Получение списка предметов заданного заказа

`GET` /order/1<br>

#### <span style="color:#12ff63">200 STATUS: OK

```json
[
  {
    "order_id": 1,
    "product_id": 1,
    "quantity": 5
  },
  {
    "order_id": 1,
    "product_id": 2,
    "quantity": 10
  }
]
```

#### <span style="color:#df0000">400 STATUS: BAD REQUEST

```json
{
  "message": "invalid request"
}
```

### Получение списка предметов заказов заданного продукта

`GET` /order/product/1<br>

#### <span style="color:#12ff63">200 STATUS: OK

```json
[
  {
    "order_id": 1,
    "product_id": 1,
    "quantity": 5
  },
  {
    "order_id": 2,
    "product_id": 1,
    "quantity": 10
  }
]
```

#### <span style="color:#df0000">400 STATUS: BAD REQUEST

```json
{
  "message": "invalid request"
}
```

### Удаление заданного предмета заказа

`DELETE` /order/item/1<br>

#### <span style="color:#12ff63">200 STATUS: OK

```json
{
  "status": "ok"
}
```

#### <span style="color:#df0000">400 STATUS: BAD REQUEST

```json
{
  "message": "order item with specified id not found"
}
```
