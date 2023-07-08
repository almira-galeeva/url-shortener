# url-shortener
Сервис, предоставляющий API по сокращению ссылок.

## API
### Метод POST, который сохраняет оригинальный URL в базе и возвращает сокращенный.
**POST** `{base_url}/shortener/short_url`
В теле запроса передается оригинальная ссылка:
```
{"original_url": "https://www.ozon.ru/"}
```

В ответе придет сокращенная ссылка:
```
{
    "shortUrl": "hjgfhjl"
}
```

### Метод GET, который принимает сокращенный URL и возвращает оригинальный.
**GET** `{base_url}/shortener/original_url/{short_url}`

В ответе придет оригинальная ссылка:
```
{
    "originalUrl": "ut sit"
}
```


