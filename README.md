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


При попытке передать в любой метод строку, не являющуюся ссылкой, получим ошибку:
```
parse "{переданная невалидная ссылка}": invalid URI for request
```
