
# yaml2json 1.0.0

Converts incoming json, xml or properties to json, xml or properties.

---
- #### Category: Tools
- #### Image: direktiv/yaml2json 
- #### License: [Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)
- #### Issue Tracking: https://github.com/direktiv-apps/yaml2json/issues
- #### URL: https://github.com/direktiv-apps/yaml2json
- #### Maintainer: [direktiv.io](https://www.direktiv.io)
---

## About yaml2json

This [yq](https://github.com/mikefarah/yq) based service can convert incoming json, xml, csv, tsv or properties data into the same formats.
The incoming data can be a string, file path or base64-encoded string. 

### Example(s)
  #### Function Configuration
  ```yaml
  functions:
  - id: yaml2json
    image: direktiv/yaml2json
    type: knative-workflow
  ```
   #### String
   ```yaml
   - id: convert
     type: action
     action:
       function: yaml2json
       input:
         data: 'hello: world'
         output-format: json
         input-type: string
   ```
   #### File
   ```yaml
   - id: set
     type: setter
     variables:
     - key: indata
       scope: instance
       mimeType: text/plain
       value: 'hello: world'
     transition: modify
  - id: modify 
     type: action
     action:
       function: yaml2json
       files:
         - key: indata
           scope: instance
       input: 
         data: indata
         output-format: json
         input-type: file
   ```
   #### Base64
   ```yaml
   - id: convert
     type: action
     action:
       function: yaml2json
       input:
         data: aGVsbG86IHdvcmxk
         output-format: xml
         input-type: base64
   ```

### Request

The service requires only the data field but can accept different data types.

#### Request Attributes
[PostParamsBody](#post-params-body)

### Response
  nice greeting
#### Reponse Types
    
  

[PostOKBody](#post-o-k-body)
#### Example Reponses
    
```json
"Hello YourName"
```

### Errors
| Type | Description
|------|---------|
| io.direktiv.command.error | Command execution failed |
| io.direktiv.output.error | Template error for output generation of the service |
| io.direktiv.ri.error | Can not create information object from request |


### Types
#### <span id="post-o-k-body"></span> postOKBody

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| output | [PostOKBodyOutput](#post-o-k-body-output)| `PostOKBodyOutput` |  | |  |  |


#### <span id="post-o-k-body-output"></span> postOKBodyOutput

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| result | [interface{}](#interface)| `interface{}` |  | |  |  |
| success | boolean| `bool` |  | |  |  |


#### <span id="post-params-body"></span> postParamsBody

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| data | string| `string` | ✓ | | Depending on the input-type this value can be a file path, base64 string or a plain string. | `aHR0cHM6Ly93d3cueW91dHViZS5jb20vd2F0Y2g/dj1kUXc0dzlXZ1hjUQ==` |
| input-format | string| `string` |  | `"yaml"`| Input format. Possible values are yaml,xml,prop. | `xml` |
| input-type | string| `string` |  | `"string"`|  | `base64` |
| output-format | string| `string` |  | `"json"`| Output format. Possible values are yaml, json, props, csv, tsv, xml. | `json` |

 
