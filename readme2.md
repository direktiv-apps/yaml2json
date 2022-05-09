


# yaml2json
Converts incoming json, xml or properties to json, xml or properties.
  

## Informations

### Version

1.0.0

## Content negotiation

### URI Schemes
  * http

### Consumes
  * application/json

### Produces
  * application/json

## All endpoints

###  operations

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| DELETE | / | [delete](#delete) |  |
| POST | / | [post](#post) |  |
  


## Paths

### <span id="delete"></span> delete (*Delete*)

```
DELETE /
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| Direktiv-ActionID | `header` | string | `string` |  |  |  | On cancel Direktiv sends a DELETE request to
the action with id in the header |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-200) | OK |  |  | [schema](#delete-200-schema) |

#### Responses


##### <span id="delete-200"></span> 200
Status: OK

###### <span id="delete-200-schema"></span> Schema

### <span id="post"></span> post (*Post*)

```
POST /
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| Direktiv-ActionID | `header` | string | `string` |  |  |  | direktiv action id is an UUID. 
For development it can be set to 'development' |
| Direktiv-TempDir | `header` | string | `string` |  |  |  | direktiv temp dir is the working directory for that request
For development it can be set to e.g. '/tmp' |
| body | `body` | [PostParamsBody](#post-params-body) | `models.PostParamsBody` | |  | | The service requires only the data field but can accept different data types. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#post-200) | OK | nice greeting |  | [schema](#post-200-schema) |
| [default](#post-default) | | generic error response | ✓ | [schema](#post-default-schema) |

#### Responses


##### <span id="post-200"></span> 200 - nice greeting
Status: OK

###### <span id="post-200-schema"></span> Schema
   
  

[PostOKBody](#post-o-k-body)

###### Examples
    
**greeting**
```json
"Hello YourName"
```

##### <span id="post-default"></span> Default Response
generic error response

###### <span id="post-default-schema"></span> Schema

  

[Error](#error)

###### Response headers

| Name | Type | Go type | Separator | Default | Description |
|------|------|---------|-----------|---------|-------------|
| Direktiv-ErrorCode | string | `string` |  |  |  |
| Direktiv-ErrorMessage | string | `string` |  |  |  |

## Models

### <span id="error"></span> error


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| errorCode | string| `string` | ✓ | |  |  |
| errorMessage | string| `string` | ✓ | |  |  |



### <span id="post-o-k-body"></span> postOKBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| output | [PostOKBodyOutput](#post-o-k-body-output)| `PostOKBodyOutput` |  | |  |  |



### <span id="post-o-k-body-output"></span> postOKBodyOutput


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| result | [interface{}](#interface)| `interface{}` |  | |  |  |
| success | boolean| `bool` |  | |  |  |



### <span id="post-params-body"></span> postParamsBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| data | string| `string` | ✓ | | Depending on the input-type this value can be a file path, base64 string or a plain string. | `aHR0cHM6Ly93d3cueW91dHViZS5jb20vd2F0Y2g/dj1kUXc0dzlXZ1hjUQ==` |
| input-format | string| `string` |  | `"yaml"`| Input format. Possible values are yaml,xml,prop. | `xml` |
| input-type | string| `string` |  | `"string"`|  | `base64` |
| output-format | string| `string` |  | `"json"`| Output format. Possible values are yaml, json, props, csv, tsv, xml. | `json` |


