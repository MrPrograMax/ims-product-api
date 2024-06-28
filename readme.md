# User Product Service

## 1. Работа с продуктом

### Добавление продукта

`POST` /product/add<br>
Ожидаемый формат входных данных:

```json
{
  "name": "kartoshka",
  "quantity": 1,
  "description": "desc",
  "category": {
    "id": 1,
    "name": "food"
  },
  "location": {
    "id": 1,
    "row": "section 18",
    "place": "place 22"
  },
  "status": {
    "id": 1,
    "name": "active"
  }
}
```

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

### Поиск продукта по имени

`GET` /product/name/kartoshka<br>

#### <span style="color:#12ff63">200 STATUS: OK

```json
{
  "id": 1,
  "name": "kartoshka",
  "quantity": 1,
  "description": "desc",
  "category": {
    "id": 1,
    "name": "food"
  },
  "location": {
    "id": 1,
    "row": "section 18",
    "place": "place 22"
  },
  "status": {
    "id": 1,
    "name": "active"
  }
}
```

#### <span style="color:#df0000">400 STATUS: BAD REQUEST

```json
{
  "message": "param name of invalid type"
}
```

### Поиск продукта по идентификатору

`GET` /product/id/1<br>

#### <span style="color:#12ff63">200 STATUS: OK

```json
{
  "id": 1,
  "name": "kartoshka",
  "quantity": 1,
  "description": "desc",
  "category": {
    "id": 1,
    "name": "food"
  },
  "location": {
    "id": 1,
    "row": "section 18",
    "place": "place 22"
  },
  "status": {
    "id": 2,
    "name": "active"
  }
}
```

#### <span style="color:#df0000">400 STATUS: BAD REQUEST

```json
{
  "message": "param id of invalid type"
}
```

### Получение всех продуктов

`GET` /product/<br>

#### <span style="color:#12ff63">200 STATUS: OK

```json
[
  {
    "id": 1,
    "name": "kartoshka",
    "quantity": 1,
    "description": "desc",
    "category": {
      "id": 1,
      "name": "food"
    },
    "location": {
      "id": 1,
      "row": "section 18",
      "place": "place 22"
    },
    "status": {
      "id": 1,
      "name": "active"
    }
  },
  {
    "id": 2,
    "name": "botinki",
    "quantity": 2,
    "description": "desc",
    "category": {
      "id": 2,
      "name": "clothes"
    },
    "location": {
      "id": 2,
      "row": "section 18",
      "place": "place 23"
    },
    "status": {
      "id": 2,
      "name": "inactive"
    }
  }
]
```

### Поиск продуктов по идентификатору категории

`GET` /product/category/1<br>

#### <span style="color:#12ff63">200 STATUS: OK

```json
[
  {
    "id": 1,
    "name": "kartoshka",
    "quantity": 1,
    "description": "desc",
    "category": {
      "id": 1,
      "name": "food"
    },
    "location": {
      "id": 1,
      "row": "section 18",
      "place": "place 22"
    },
    "status": {
      "id": 1,
      "name": "active"
    }
  }
]
```

#### <span style="color:#df0000">400 STATUS: BAD REQUEST

```json
{
  "message": "param id of invalid type"
}
```

### Поиск продуктов по имени категории

`GET` /product/category/clothes<br>

#### <span style="color:#12ff63">200 STATUS: OK

```json
[
  {
    "id": 2,
    "name": "botinki",
    "quantity": 2,
    "description": "desc",
    "category": {
      "id": 2,
      "name": "clothes"
    },
    "location": {
      "id": 2,
      "row": "section 18",
      "place": "place 23"
    },
    "status": {
      "id": 2,
      "name": "inactive"
    }
  }
]
```

#### <span style="color:#df0000">400 STATUS: BAD REQUEST

```json
{
  "message": "param name of invalid type"
}
```

### Поиск продуктов по идентификатору локации

`GET` /product/loc/1<br>

#### <span style="color:#12ff63">200 STATUS: OK

```json
[
  {
    "id": 1,
    "name": "kartoshka",
    "quantity": 1,
    "description": "desc",
    "category": {
      "id": 1,
      "name": "food"
    },
    "location": {
      "id": 1,
      "row": "section 18",
      "place": "place 22"
    },
    "status": {
      "id": 1,
      "name": "active"
    }
  }
]
```

#### <span style="color:#df0000">400 STATUS: BAD REQUEST

```json
{
  "message": "param id of invalid type"
}
```

### Поиск продуктов по локации

`GET` /product/loc<br>
Ожидаемый формат входных данных:

```json
{
  "id": 2,
  "row": "section 18",
  "place": "place 23"
}
```

#### <span style="color:#12ff63">200 STATUS: OK

```json
[
  {
    "id": 2,
    "name": "botinki",
    "quantity": 2,
    "description": "desc",
    "category": {
      "id": 2,
      "name": "clothes"
    },
    "location": {
      "id": 2,
      "row": "section 18",
      "place": "place 23"
    },
    "status": {
      "id": 2,
      "name": "inactive"
    }
  }
]
```

#### <span style="color:#df0000">400 STATUS: BAD REQUEST

```json
{
  "message": "param row of invalid type"
}
```

#### <span style="color:#df0000">400 STATUS: BAD REQUEST

```json
{
  "message": "param place of invalid type"
}
```

### Поиск продуктов по имени статуса

`GET` /product/status/active<br>

#### <span style="color:#12ff63">200 STATUS: OK

```json
[
  {
    "id": 1,
    "name": "kartoshka",
    "quantity": 1,
    "description": "desc",
    "category": {
      "id": 1,
      "name": "food"
    },
    "location": {
      "id": 1,
      "row": "section 18",
      "place": "place 22"
    },
    "status": {
      "id": 1,
      "name": "active"
    }
  }
]
```

#### <span style="color:#df0000">400 STATUS: BAD REQUEST

```json
{
  "message": "param name of invalid type"
}
```

### Поиск продуктов по идентификатору статуса

`GET` /product/status/2<br>

#### <span style="color:#12ff63">200 STATUS: OK

```json
[
  {
    "id": 2,
    "name": "botinki",
    "quantity": 2,
    "description": "desc",
    "category": {
      "id": 2,
      "name": "clothes"
    },
    "location": {
      "id": 2,
      "row": "section 18",
      "place": "place 23"
    },
    "status": {
      "id": 2,
      "name": "inactive"
    }
  }
]
```

#### <span style="color:#df0000">400 STATUS: BAD REQUEST

```json
{
  "message": "param id of invalid type"
}
```

### Изменение продукта

`PUT` /product/1<br>
Ожидаемый формат входных данных:

```json
{
  "name": "kartoshka",
  "quantity": 1,
  "description": "Простая картошка",
  "category_id": 1,
  "location_id": 2,
  "status_id": 1
}
```

Среди входных данных могут быть перечислены не все входные параметры,
но не пустой JSON, что повлечет за собой ошибку 400 Bad Request

#### <span style="color:#12ff63">200 STATUS: OK

```json
{
  "status": "ok"
}
```

#### <span style="color:#df0000">400 STATUS: BAD REQUEST

```json
{
  "message": "updating entity is empty"
}
```

### Удаление продукта

`DELETE` /product/1<br>

#### <span style="color:#12ff63">200 STATUS: OK

```json
{
  "status": "ok"
}
```

#### <span style="color:#df0000">400 STATUS: BAD REQUEST

```json
{
  "message": "product with specified id not found"
}
```

## 2. Работа с категорией

### Получение списка категорий

`GET` /product/category/<br>

#### <span style="color:#12ff63">200 STATUS: OK

```json
[
  {
    "id": 1,
    "name": "food"
  },
  {
    "id": 2,
    "name": "clothes"
  }
]
```

### Получение категории по имени

`GET` /product/category/name/food<br>

#### <span style="color:#12ff63">200 STATUS: OK

```json
{
  "id": 1,
  "name": "food"
}
```

#### <span style="color:#df0000">400 STATUS: BAD REQUEST

```json
{
  "message": "param name of invalid type"
}
```

### Получение категории по идентификатору

`GET` /product/category/id/2<br>

#### <span style="color:#12ff63">200 STATUS: OK

```json
{
  "id": 2,
  "name": "clothes"
}
```

#### <span style="color:#df0000">400 STATUS: BAD REQUEST

```json
{
  "message": "param id of invalid type"
}
```

## 3. Работа с локацией

### Добавление локации

`POST` /product/location/add<br>
Ожидаемый формат входных данных:

```json
{
  "row": "section 19",
  "place": "place 23"
}
```

#### <span style="color:#12ff63">200 STATUS: OK

```json
{
  "id": 3
}
```

#### <span style="color:#df0000">400 STATUS: BAD REQUEST

```json
{
  "message": "invalid request"
}
```

### Поиск локации по ряду и месту

`GET` /product/location/row/section%2018/place/place%2022<br>

#### <span style="color:#12ff63">200 STATUS: OK

```json
{
  "id": 1,
  "row": "section 18",
  "place": "place 22"
}
```

#### <span style="color:#df0000">400 STATUS: BAD REQUEST

```json
{
  "message": "param name of invalid type"
}
```

### Поиск локации по идентификатору локации

`GET` /product/location/id/2<br>

#### <span style="color:#12ff63">200 STATUS: OK

```json
{
  "id": 2,
  "row": "section 18",
  "place": "place 23"
}
```

#### <span style="color:#df0000">400 STATUS: BAD REQUEST

```json
{
  "message": "param id of invalid type"
}
```

### Получение списка всех локаций

`GET` /product/location/<br>

#### <span style="color:#12ff63">200 STATUS: OK

```json
[
  {
    "id": 1,
    "row": "section 18",
    "place": "place 22"
  },
  {
    "id": 2,
    "row": "section 18",
    "place": "place 23"
  }
]
```

### Получение списка локаций по ряду

`GET` /product/location/row/section%2018<br>

#### <span style="color:#12ff63">200 STATUS: OK

```json
[
  {
    "id": 1,
    "row": "section 18",
    "place": "place 22"
  },
  {
    "id": 2,
    "row": "section 18",
    "place": "place 23"
  }
]
```

#### <span style="color:#df0000">400 STATUS: BAD REQUEST

```json
{
  "message": "param id of invalid type"
}
```

### Удаление локации

`DELETE` /product/location/1<br>

#### <span style="color:#12ff63">200 STATUS: OK

```json
{
  "status": "ok"
}
```

#### <span style="color:#df0000">400 STATUS: BAD REQUEST

```json
{
  "message": "location with specified id not found"
}
```

## 4. Работа со статусом продукта

### Получение списка статусов продукта

`GET` /product/status/<br>

#### <span style="color:#12ff63">200 STATUS: OK

```json
[
  {
    "id": 1,
    "name": "active"
  },
  {
    "id": 2,
    "name": "inactive"
  }
]
```

### Получение статуса продукта по имени

`GET` /product/status/name/inactive<br>

#### <span style="color:#12ff63">200 STATUS: OK

```json
{
  "id": 2,
  "name": "inactive"
}
```

#### <span style="color:#df0000">400 STATUS: BAD REQUEST

```json
{
  "message": "param name of invalid type"
}
```

### Получение статуса продукта по идентификатору

`GET` /product/status/id/1<br>

#### <span style="color:#12ff63">200 STATUS: OK

```json
{
  "id": 1,
  "name": "active"
}
```

#### <span style="color:#df0000">400 STATUS: BAD REQUEST

```json
{
  "message": "param id of invalid type"
}
```
