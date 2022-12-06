# Json configuration file for the project

## 1. File Description

| <h3> Configure File | <h3> Description |
|:-------------:|:-----------:|
| new1.json | route sample api urls. <br> Result : don't passed Header. |
| new2.json | Add Header 'Authorization' <br> Result : Passed Header. (JWT Authorized)|
| new3.json | Using Https url -self signed certificate <br> Result : Pass all urls |
| new4.json | 'new3.json' Add header 'Cookie' <br> Result : don't passed cookie. Backend not recognizing Cookie header |
| new5.json | Add all header pass and change encoding 'no-op' <br> Result : Passed 'Cookie' header. |
| new6.json | Merging reponse service A,B <br> Result : response merged and field name at config's 'group' field |

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
// original backend response 
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
```
| <h3>Option | <h3>Description | <h3>example</h3> |
|:-------------:|:-----------:|:-------|
| group | add reponse's  root field name | // backend1 -> "group":"backend1" </br> // backend2 -> "group":"backend2" </br>{ "backend1" : { "service":"A"  ,"message" : "ServiceA" }, </br> "backend2" : { "service": "B", "message" : "ServiceB"} } |
| allow | within field name | // backend1 -> "allow" : [ "service" ] </br> // backend2 -> "allow" : [ "message" ]</br> { "backend1" : { "service":"A"   }, "backend2" : { "message" : "ServiceB"}  } |
| deny | without field name | // backend1 -> "deny" : [ "message" ] </br> // backend2 -> "deny" : [ "service" ]</br> { "backend1" : { "service":"A"   }, "backend2" : { "message" : "ServiceB"} } |
| mapping | replace field name | // backend1 -> "mapping" : { "service":"service_name" } </br> // backend2 -> "mapping" : { "message" : "contents" } </br> {"backend1":{"service_name":"A"},"backend2":{"contents":"ServiceB"}} |
