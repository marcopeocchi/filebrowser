# filebrowser

A lightweight minimalistic filebrowser.

![image.png](https://i.postimg.cc/0y364LvB/image.png)

## Installation
```shell
# download latest release for your arch
# edit db.json

# -r  -> root directory
# -p  -> port to listen at
# -db -> database file path (user must have read permissions)

filebrowser --db /var/filebrowser/db.json -r /mnt/downloads -p 8080
```

This filebrowser is intended to be used behind a reverse proxy like Caddy and served through HTTPS (if you're a security maniac like me). 


## JSONDB file database

```json5
// Default entry for db.json
// uid can be anything
// password is a bcrypt hashed password,
// the following example is bcrypt hashed "adminadmin"
// role: 0 -> admin; 1 -> user
[
  {
    "uid": "1",
    "username": "admin",
    "password": "$2a$12$ER0Jhb1BRriXnoD7DR39WO388vpWO/4DYPy06G58JsLXmZKkezr3i",
    "role": 0
  }
]
```

Golang is king.
