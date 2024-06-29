# Supply Service

- ### [1. Поставка](#1-поставка)
    - [Добавление поставки](#добавление-поставки)
    - [Получение всех поставок](#получение-всех-поставок)
    - [Удаление поставки](#удаление-поставки)
- ### [2. Предметы поставки](#2-предметы-поставки)
    - [Добавление предмета поставки](#добавление-предмета-поставки)
    - [Получение списка предметов поставки по supply id](#получение-списка-предметов-заданной-поставки)
    - [Получение списка предметов поставки по product id](#получение-списка-предметов-поставок-заданного-продукта)
    - [Удаление предмета поставки](#удаление-заданного-предмета-поставки)

## 1. Поставка

### Добавление поставки

`POST` /supply<br>

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

### Получение всех поставок

`GET` /supply<br>

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

### Удаление поставки

`DELETE` /supply/1<br>

#### <span style="color:#12ff63">200 STATUS: OK

```json
{
  "status": "ok"
}
```

#### <span style="color:#df0000">400 STATUS: BAD REQUEST

```json
{
  "message": "supply with specified id not found"
}
```

## 2. Предметы поставки

### Добавление предмета поставки

`POST` /supply/item<br>
Ожидаемый формат входных данных:

```json
{
  "supply_id": 1,
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

### Получение списка предметов заданной поставки

`GET` /supply/1<br>

#### <span style="color:#12ff63">200 STATUS: OK

```json
[
  {
    "supply_id": 1,
    "product_id": 1,
    "quantity": 5
  },
  {
    "supply_id": 1,
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

### Получение списка предметов поставок заданного продукта

`GET` /supply/product/1<br>

#### <span style="color:#12ff63">200 STATUS: OK

```json
[
  {
    "supply_id": 1,
    "product_id": 1,
    "quantity": 5
  },
  {
    "supply_id": 2,
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

### Удаление заданного предмета поставки

`DELETE` /supply/item/1<br>

#### <span style="color:#12ff63">200 STATUS: OK

```json
{
  "status": "ok"
}
```

#### <span style="color:#df0000">400 STATUS: BAD REQUEST

```json
{
  "message": "supply item with specified id not found"
}
```
