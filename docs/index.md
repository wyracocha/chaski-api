## Welcome to Chaski-Api

You can use it api for mocking your http request sending headers. You can send simple strings or json content using the key **Payload** in the header

```sh
curl  -H 'Payload: {"key": "value"}' https://chaski-api.wyracocha.com
```
The response is returning in next format:
```sh
{
  "body": "{\"key\": \"value\"}"
}
```
## Roadmap
- Support for **api keys**
- Add **rate limit**
