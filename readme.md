# User Product Service

## Указатель

- ### [Запуск](#Запуск)
- ### [Сервис продуктов](docs/readme/product.md)
- ### [Сервис поставок](docs/readme/supply.md)
- ### [Сервис заказов](docs/readme/order.md)

## Запуск
Подгрузка недостающих библиотек
```cmd
go mod tidy
```
Создайте файл .env и добавьте пароль:
```yaml
DB_PASSWORD=qwerty
```
База данных
```cmd
docker run --name=ims-product-db -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 --rm postgres
```
Миграция
```cmd
make migrate
```
или
```cmd
migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5432/postgres?sslmode=disable' up
```
Приложение
```
go run cmd/main.go
```
