# task-schema-test

Test different ways to represents general task data.

Server exposes 2 endpoints:
- `/raw-message`: return task with `Data` field defined using `json.RawMessage`
- `/any`: return task with `Data` field defined using `any`

Client makes requests to above 2 endpoints and for each endpoint, it will try to unmarshal response into task with `json.RawMessage` field and task with `any` field

To run: on 1 terminal `task run-server`, on another terminal `task run-client`

Output:

```
endpoint /raw-message response body as string: {"data":{"name":"custom-resource"}}
endpoint /raw-message - unmarshalled into raw message, data type: json.RawMessage, data: {Data:[123 34 110 97 109 101 34 58 34 99 117 115 116 111 109 45 114 101 115 111 117 114 99 101 34 125]}
endpoint /raw-message - unmarshalled into any, data type: map[string]interface {}, data: {Data:map[name:custom-resource]}
endpoint /any response body as string: {"data":{"name":"custom-resource"}}
endpoint /any - unmarshalled into raw message, data type: json.RawMessage, data: {Data:[123 34 110 97 109 101 34 58 34 99 117 115 116 111 109 45 114 101 115 111 117 114 99 101 34 125]}
endpoint /any - unmarshalled into any, data type: map[string]interface {}, data: {Data:map[name:custom-resource]}
```
