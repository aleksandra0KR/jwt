# Часть сервиса аутентификации.

Два REST маршрута:

- Первый маршрут выдает пару Access, Refresh токенов для пользователя с идентификатором (GUID) указанным в параметре запроса
- Второй маршрут выполняет Refresh операцию на пару Access, Refresh токенов

# Запуск в Docker
Склонировать проект с гита

```
git clone https://github.com/aleksandra0KR/jwt
```
Перейти в директорию проекта
```
cd jwt
```
Забилдить
```
docker compose build
```
Запустить:
```
docker compose up
```
---
# Запустить без Docker

```
git clone https://github.com/aleksandra0KR/jwt
```
Перейти в директорию проекта
```
cd jwt
```
Запустить
```
go run cmd/main.go
```

### В файле .env можно поменять на нужные вам параметры : postgres или in-memory



# Пример работы
- Первый маршрут Метод POST
  ```
  localhost:8080/api/auth/123
  ```
  
  Тело:
  ```
  {
      "IP"    :     "456789"
  }
  ```
  ![](https://github.com/aleksandra0KR/jwt/tree/main/examples/1.png?raw=true)- 
- Второй маршрут Метод POST
  ```
  localhost:8080/api/refreshToken/123
  ```

  Тело:
  ```
  {
      "IP"    :     "456789"
  }
  ```
  ![](https://github.com/aleksandra0KR/jwt/tree/main/examples/2.png?raw=true)