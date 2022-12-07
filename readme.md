# Json configuration file for the project

## 1. File Description

| <h3> Configure File | <h3> Description |
|:-------------:|:-----------:|
| new1.json | route sample api urls. <br> Result : don't passed Header. |
| new2.json | Add Header 'Authorization' <br> Result : Passed Header. (JWT Authorized)|
| new3.json | Using Https url -self signed certificate <br> Result : Pass all urls |
| new4.json | 'new3.json' Add header 'Cookie' <br> Result : don't passed cookie. Backend not recognizing Cookie header |
| new5.json | Add all header pass and change encoding 'no-op' <br> Result : Passed 'Cookie' header. |
| new6.json | manipulation reponse </br> [ move [Response manipulation] ](#response-manipulation-new6json) |
| new7.json | Http error status reponse |
---

## 2. Endpoint urls

| <h3> Url | <h3> Description |
|:-------------:|:-----------:|
| /api/v1/authorize | Create and return JWT token in message, cookie |
| /api/v1/authorize/verify | Check 'Authorization' header |
| /api/v1/authorize/verfiycookie | Check 'Cookie' header |

---

## 3. Krakend Routing urls
| <h3>Krakend  | >> | <h3>Backend | <h3>description |
|:-------------:|:--:|:-----------:|:-----------:|
| /create | >> | /api/v1/authorize | create and return jwt token |
| /verify | >> | /api/v1/authorize/verify | send 'Authorization' header |
| /verify/cookie | >> | /api/v1/authorize/verifycookie | send 'Cookie' header |

---
## 4.Merge sample
| <h3>Krakend  | >> | <h3>Backend | <h3>description |
|:-------------:|:--:|:-----------:|:-----------:|
| /ab | >> | serviceA / serviceB | response merge serviceA and serviceB |

---
</br></br>

# Response manipulation (new6.json)

``` json
// backend response 
// backend1 

{ 
    "service": "A",
    "message": "ServiceA"
}

// backend2
{ 
    "service": "B",
    "message": "ServiceB"
}
// backend3
{ 
    "service": "C",
    "message": "ServiceC"
    "data": {
        "name": "sam",
        "age": 30
    }
}
```
| <h3>Option | <h3>Description | <h3>example</h3> |
|:-------------:|:-----------:|:-------|
| group | add reponse's  root field name | // backend1 -> "group":"backend1" </br> // backend2 -> "group":"backend2" </br>{ "backend1" : { "service":"A"  ,"message" : "ServiceA" }, </br> "backend2" : { "service": "B", "message" : "ServiceB"} } |
| allow | within field name | // backend1 -> "allow" : [ "service" ] </br> // backend2 -> "allow" : [ "message" ]</br> { "backend1" : { "service":"A"   }, "backend2" : { "message" : "ServiceB"}  } |
| deny | without field name | // backend1 -> "deny" : [ "message" ] </br> // backend2 -> "deny" : [ "service" ]</br> { "backend1" : { "service":"A"   }, "backend2" : { "message" : "ServiceB"} } |
| mapping | replace field name | // backend1 -> "mapping" : { "service":"service_name" } </br> // backend2 -> "mapping" : { "message" : "contents" } </br> {"backend1":{"service_name":"A"},"backend2":{"contents":"ServiceB"}} |
| target | extract field | // backend3 -> "target" : "data" </br> {"age":30,"name":"sam"} |


# Flexible Krakend Config
## ref : Contents directory : flexible-config

| <h3> Directory | <h3>File | <h3>Description |
|:--------------:|:-------------:|:-----------:|
| settings |  |sub configure files(backend, endpoint) </br> (used in krakend configuration) |
| settings | service.json | krakend daemon configure (port, encoding, tls settigs, etc ...) |
| settings | single_endpoint.json | endpoint settings (normal endpoint) |
| settings | merge_endpoint.json | endpoint settings (2 reponse merged) |
| partials |  |sub configure files(extra-config) </br> (used in krakend configuration) |

--- 
## Krakend syntax check scripts : [check_syntax.sh](./flexible-config/check_syntax.sh)
### Usage : ./check_syntax.sh <output_json_file> <krakend_json_file>
  
--- 

## Environment variables - reference: [Docs krakend -flexible-config](https://www.krakend.io/docs/flexible-config/)

`FC_ENABLE=1` to let KrakenD know that you are using Flexible Configuration. You can use 1 or any other value (but 0 won’t disable it!). The file passed with the -c flag is the base template.

`FC_SETTINGS=dirname`: The path to the directory with all the settings files.


`FC_PARTIALS=dirname`: The path to the directory with the partial files included in the configuration file. Partial files DON’T EVALUATE, they are only inserted in the placeholder.


`FC_TEMPLATES=dirname`: The path to the directory with the sub-templates included in the configuration file. These are evaluated using the Go templating system.


`FC_OUT`: For debugging purposes, saves the resulting configuration of processing the flexible configuration in the given filename. Otherwise, the final file is not visible.
