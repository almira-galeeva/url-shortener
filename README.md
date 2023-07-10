# url-shortener
Сервис, предоставляющий API по сокращению ссылок.

Чтобы поднять сервис локально, необходимо 
- склонировать репозиторий
- ввести команду `make run`

## API
### Метод POST, который сохраняет оригинальный URL в базе и возвращает сокращенный.
##### HTTP
**POST**: `{base_url}/shortener/short_url`
В теле запроса передается оригинальная ссылка:
```
{"original_url": "https://github.com/almira-galeeva/url-shortener"}
```

В ответе придет сокращенная ссылка:
```
{
    "shortUrl": "https://shorturl.com/LTiekH83dR"
}
```
##### GRPC
`Shortener/GetShortUrl`

В Message передается оригинальная ссылка ссылка
```
{
    "original_url": "https://github.com/almira-galeeva/url-shortener"
}
```
В ответе придется сокращенная ссылка:
```
{
    "short_url": "https://shorturl.com/LTiekH83dR"
}
```

### Метод GET, который принимает сокращенный URL и возвращает оригинальный.
##### HTTP
Передаем в ручку хэш, сформированный в методе POST
**GET**: `{base_url}/shortener/original_url/{short_url_hash}`

В ответе придет оригинальная ссылка:
```
{
    "originalUrl": "https://github.com/almira-galeeva/url-shortener"
}
```
##### GRPC
`Shortener/GetOriginalUrl`

В Message передается хэш, сформированный в методе POST
```
{
    "short_url": "LTiekH83dR"
}
```
В ответе придется оригинальная ссылка:
```
{
    "original_url": "https://github.com/almira-galeeva/url-shortener"
}
```

### Проверки

1. При попытке передать в любой метод строку, не являющуюся ссылкой, получаем ошибку:
```
parse "{переданная невалидная ссылка}": invalid URI for request
```

2. При попытке передать в метод POST ссылку, которая уже есть в базе данных/памяти, получаем ошибку:
```
Url {переданная ссылка} already exists in db
```
3. При попытке передать в метод GET ссылку, которой нет в базе данных/памяти, получаем ошибку
```
no rows in result set
```
