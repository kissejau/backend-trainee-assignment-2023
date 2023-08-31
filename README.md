# [Тестовое задание (clickable)](https://github.com/avito-tech/backend-trainee-assignment-2023)

## Запуск:

```
docker compose up
```

## Задача:

Требуется реализовать сервис, хранящий пользователя и сегменты, в которых он состоит (создание, изменение, удаление сегментов, а также добавление и удаление пользователей в сегмент)


## Основное задание (минимум):

- [x] Метод создания сегмента. Принимает slug (название) сегмента. 
- [x] Метод удаления сегмента. Принимает slug (название) сегмента. 
- [x] Метод добавления пользователя в сегмент. Принимает список slug (названий) сегментов которые нужно добавить пользователю, список slug (названий) сегментов которые нужно удалить у пользователя, id пользователя.
- [x] Метод получения активных сегментов пользователя. Принимает на вход id пользователя.


## Опциональные задания

- [x] реализовать сохранение истории попадания/выбывания пользователя из сегмента с возможностью получения отчета по пользователю за определенный период. На вход: год-месяц. На выходе ссылка на CSV файл.

- [x] реализовать возможность задавать TTL (время автоматического удаления пользователя из сегмента)

- [ ] в методе создания сегмента, добавить опцию указания процента пользователей, которые будут попадать в сегмент автоматически. В методе получения сегментов пользователя, добавленный сегмент должен отдаваться у заданного процента пользователей.

## Взаимодействие с сегментами

### POST /api/v1/segment

#### Создание сегмента
- Тело запроса
  - `slug` - название сегмента
- Тело ответа (code 200)
  - `slug` - название сегмента

Запрос:
```
curl --location 'localhost:8080/api/v1/segment' \
--header 'Content-Type: application/json' \
--data '{
    "slug": "AVITO_VOICE_MESSAGES"
}'
```
Ответ:
```
{
	"slug": "AVITO_VOICE_MESSAGES"
}
```

### GET /api/v1/segments

#### Получение всех сегментов

Запрос:
```
curl --location 'localhost:8080/api/v1/segments' \
--header 'id: 1'
```
Ответ:
```
[
    {
        "id": "1",
        "slug": "AVITO_VOICE_MESSAGES"
    }
]
```

### GET /api/v1/segment

#### Получение сегмента по его названию
- Заголовки
  - `slug` - название сегмента
- Тело ответа (code 200)
  - `id` - идентификатор сегмента
  - `slug` - название сегмента

Запрос:
```
curl --location 'localhost:8080/api/v1/segment' \
--header 'slug: AVITO_VOICE_MESSAGES'
```
Ответ:
```
{
    "id": "1",
    "slug": "AVITO_VOICE_MESSAGES"
}
```

### PATCH /api/v1/segment

#### Обновление сегмента

- Тело запроса
  - `id` - идентификатор сегмента
  - `slug` - название сегмента
- Тело ответа (code 200)
  - `id` - идентификатор сегмента
  - `slug` - название сегмента

Запрос:
```
  curl --location --request PATCH 'localhost:8080/api/v1/segment' \
--header 'Content-Type: application/json' \
--data '{
    "id": "1",
    "slug": "AVITO_DISCOUNT_30"
}'
```
Ответ:
```
{
    "id": "1",
    "slug": "AVITO_DISCOUNT_30"
}
```

### DELETE /api/v1/segment

#### Обновление сегмента

- Заголовки
  - `id` - идентификатор сегмента
- Тело ответа (code 200)
  - `id` - идентификатор сегмента
  - `slug` - название сегмента

Запрос:
```
curl --location --request DELETE 'localhost:8080/api/v1/segment' \
--header 'id: 1'
```

## Взаимодействие с пользователями

### POST /api/v1/user

#### Создание пользователя

- Тело запроса
  - `name` - имя пользователя
- Тело ответа (code 200)
  - `name` - имя пользователя

Запрос:
```
curl --location 'localhost:8080/api/v1/user' \
--header 'Content-Type: application/json' \
--data '{
    "name": "kissejau"
}'
```
Ответ:
```
{
    "name": "kissejau"
}
```

### GET /api/v1/users

#### Получение всех пользователей

- Тело ответа (code 200)
  - `Array`
    - `id` - идентификатор пользователя
    - `name` - имя пользователя

Запрос:
```
curl --location 'localhost:8080/api/v1/users'
```
Ответ:
```
[
    {
        "id": "1",
        "name": "kissejau"
    }
]
```

### GET /api/v1/user

#### Получение пользователя по id

- Заголовки
  - `id` - инедтификатор пользователя
- Тело ответа (code 200)
  - `id` - инедтификатор пользователя
  - `name` - имя пользователя

Запрос:
```
curl --location 'localhost:8080/api/v1/user' \
--header 'id: 1'
```
Ответ:
```
{
    "id": "1",
    "name": "kissejau"
}
```

### UPDATE /api/v1/user

#### Обновление пользователя

- Тело запроса
  - `id` - инедтификатор пользователя
  - `name` - имя пользователя
- Тело ответа (code 200)
  - `id` - инедтификатор пользователя
  - `name` - имя пользователя

Запрос:
```
  curl --location --request PATCH 'localhost:8080/api/v1/user' \
--header 'Content-Type: application/json' \
--data '{
    "id": "1",
    "name": "kissejauxgod"
}'
```
Ответ:
```
{
    "id": "1",
    "name": "kissejauxgod"
}
```

### PATCH /api/v1/user

#### Удаление пользователя

- Заголовки
  - `id` - инедтификатор пользователя
- Тело ответа (code 200)

Запрос:
```
curl --location --request DELETE 'localhost:8080/api/v1/user' \
--header 'id: 1'
```

## Взаимодействие с сегментами пользователя

### POST /api/v1/user/segments

#### Добавление пользователя в сегмент или удаление из сегмента

- Тело запроса
  - `up_slugs` - `Array` `(OPTIONAL)`
    - `slug` - название сегмента
    - `deadline` - время удаления(истечения) сегмента `(OPTIONAL)`
  - `down_slugs` - `Array` `(OPTIONAL)`
    - `slug` - название сегмента
    - `deadline` - время удаления(истечения) сегмента `(OPTIONAL)`
  - `user_id` - идентификатор пользователя
- Тело ответа (code 200)

Запрос:
```
curl --location 'localhost:8080/api/v1/user/segments' \
--header 'Content-Type: application/json' \
--data '{
    "up_slugs": [{"slug": "AVITO_VOICE_MESSAGE", "deadline": "2024-08-29 12:59"}, {"slug": "AVITO_PERFORMANCE_VAS", "deadline": "2023-08-01 12:00"}],
    "user_id": "2"
}'
```

### GET /api/v1/user/segments

#### Получение сегментов пользователя (активных и неактивных)

- Заголовки
  - `id` - индентификатор пользователя
- Тело ответа (code 200)
  - `Array`
    - `id` - идентификатор пользователя
    - `slug` - название сегмента
    - `is_active` - флаг активности сегмента (истек deadline или нет)

Запрос:
```
curl --location 'localhost:8080/api/v1/user/segments' \
--header 'id: 2'
```
Ответ:
```
[
    {
        "id": "1",
        "slug": "AVITO_VOICE_MESSAGE",
        "is_active": true
    },
    {
        "id": "3",
        "slug": "AVITO_PERFORMANCE_VAS",
        "is_active": false
    }
]
```

## Взаимодействие с историей добавления и удаления пользователей из сегментов

### GET /api/v1/logs

#### Получение ссылки на файл csv формата с историей добавления и удаления пользователей из сегментов по заданной дате (год, месяц)

- Тело запроса
  - `year` - год для фильтрации запросов
  - `month` - месяц для фильтрации запросов
- Тело ответа
  - `link` - ссылка на csv файл

Запрос:
```
curl --location --request GET 'localhost:8080/api/v1/logs' \
--header 'Content-Type: application/json' \
--data '{
    "year": "2023",
    "month": "08"
}'
```
Ответ:
```
{
    "link": "http://localhost:8080/api/v1/log/?log=20230831201439831"
}
```

### GET /api/v1/log/{log}

#### Получение csv файла с историей добавления и удаления пользователей из сегментов

- Параметры
  - `log` - идентификатор(имя) лога
- Тип содержимого ответа (code 200)
  - `text/csv`

Запрос:
```
curl --location 'http://localhost:8080/api/v1/log/?log=20230831201439831'
```
Ответ:
```
идентификатор пользователя 1,AVITO_VOICE_MESSAGE,операция Add,2023-08-31 19:55:17.618985 +0000 +0000
идентификатор пользователя 2,AVITO_PERFORMANCE_VAS,операция Add,2023-08-31 19:55:17.633362 +0000 +0000

```
