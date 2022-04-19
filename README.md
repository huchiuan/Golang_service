


# 系統說明：
## Golang Service
+ 使用 gorilla/mux 開發 RESTful service.
+ 使用 gorm 進行資料庫存取.
+ 使用 JWT 獲取 Token 做驗證.
+ 使用 Docker 運行 Golang Server 與 PostgreSQL .
## 使用流程
+ 新增使用者的 API 與 使用者登入的API 不需要事先驗證 JWT Token.
+ 其餘所有的 API 都必須先進行使用者登入後，拿到 JWT Token 才能使用.
<br></br>

# 使用之前：
## 在使用之前必須使用 makefile 運行啟動指令。
+ 使用 Docker 運行 Golang Server 與 PostgreSQ，於 local 建立兩個服務的 Container。

<br></br>
## 請在終端機輸入 :
```terminal 
make RUN
```

### 並確認 Docker 中是否包含以下服務 :
+ Golang_service PORT : 5000
+ DB_service PORT : 5432
  
### 如服務皆已建立，即可測試服務。

 
<br></br>



# API 介紹 :

##  HTTP GET
+ 

```zsh

//Create功能
curl -X POST -k http://localhost:8090/Create -d '{"table": "kkbox","key":"edison","data":"ma"}'

//Get功能
curl -X POST -k http://localhost:8090/Get -d '{"table": "kkbox","key":"edison","data":"ma"}'

//Update功能
curl -X POST -k http://localhost:8090/Update -d '{"table": "kkbox","key":"edison","data":"ma"}'

//Delete功能
curl -X POST -k http://localhost:8090/Delete -d '{"table": "kkbox","key":"edison"}'

```