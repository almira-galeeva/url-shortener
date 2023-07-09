# url-shortener
Сервис, предоставляющий API по сокращению ссылок.

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
**GET**: `{base_url}/shortener/original_url/{short_url}`

В ответе придет оригинальная ссылка:
```
{
    "originalUrl": "https://github.com/almira-galeeva/url-shortener"
}
```
##### GRPC
`Shortener/GetOriginalUrl`

В Message передается сокращенная ссылка
```
{
    "short_url": "https://shorturl.com/LTiekH83dR"
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
4. При попытке передать в метод GET ссылку, которая начинается не с `https://shorturl.com/` (параметр `url_prefix`, который задается в конфиге `config/config.json`), получаем ошибку:
```
Short url should start with https://shorturl.com/
```
