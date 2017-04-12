# goiris
---
#### Go version: go1.7.1 darwin/amd64

This project is developed for implementation of mobile APIs. This project uses Golang and its fastest Web framework called iris.

![Gopher](http://kingofwallpapers.com/gopher/gopher-020.jpg)

## API Document
|          Description           |                        Endpoint                  |
|--------------------------------|--------------------------------------------------|
| To get the list of postoffices | http://127.0.0.1:8080/api/v1/postoffices/        |
| To fetch a single postoffice   | http://127.0.0.1:8080/api/v1/postoffices/95221   |

## (1) Request
http://127.0.0.1:8080/api/v1/postoffices/ 

### Response
<pre>
[
  {
    "pincode": 95221,
    "city": "Bilaspur",
    "state": "Himachal Pradesh",
    "district": "Bilaspur"
  },
  {
    "pincode": 110001,
    "city": "New Delhi",
    "state": "Delhi",
    "district": "New Delhi"
  },
  {
    "pincode": 110002,
    "city": "New Delhi",
    "state": "Delhi",
    "district": "Central Delhi"
  },
]
</pre>
---
### (2) Request
http://127.0.0.1:8080/api/v1/postoffices/95221

### Response
<pre>
{
  "pincode": 95221,
  "city": "Bilaspur",
  "state": "Himachal Pradesh",
  "district": "Bilaspur"
}
</pre>
