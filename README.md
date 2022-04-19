


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

##  由於 TOKEN 使用於 COOKIE 中，建議使用 POSTMAN　測試。
<br></br>

##  常用資料介紹 :
+ acct : 使用者帳號
+ pwd : 使用者密碼
+ fullname : 使用者名稱

<br></br>

#  HTTP POST :
## 建立一個使用者  (commit 4)

```zsh
// API
http://localhost:5000/create-the-user
// JSON
{"acct":"jason123","pwd":"0000","fullname":"jason"} 
```
### POSTMAN DEMO (commit 4)
![image](https://upload.cc/i1/2022/04/19/04miPR.png)
<br></br>

## 登入使用者  (commit 5)

```zsh
// API
http://localhost:5000/sign-in
// JSON
{"acct":"jason123","pwd":"0000" }
```
### POSTMAN DEMO (commit 5)
![image](https://upload.cc/i1/2022/04/19/1TEKFZ.png)
<br></br>

## 使用 fullname 搜尋使用者  (commit 2)
```zsh
// API
http://localhost:5000/search-an-user
// JSON
{"fullname":"jason"}
```
### POSTMAN DEMO  (commit 2)
![image](https://upload.cc/i1/2022/04/19/cD6RFa.png)
<br></br>


## 變更使用者密碼 (commit 7)
```zsh
// API
http://localhost:5000/update-the-user
// JSON
{"acct":"jason123","pwd":"66666" }
```
### POSTMAN DEMO  (commit 7)
![image](https://upload.cc/i1/2022/04/19/gCi2OF.png)
<br></br>

## 變更使用者名稱 (commit 9)
```zsh
// API
http://localhost:5000/update-the-user-fullname
// JSON
{"acct":"jason123","fullname":"6666666666666666666" }
```
### POSTMAN DEMO  (commit 9)
![image](https://upload.cc/i1/2022/04/19/zGlEek.png)
<br></br>

## 刪除使用者 (commit 6)
```zsh
// API
http://localhost:5000/delete-the-user
// JSON
{"acct":"jason123" }
```
### POSTMAN DEMO  (commit 6)
![image](https://upload.cc/i1/2022/04/19/R6z8Po.png)
<br></br>


#  HTTP GET :
## 列出所有使用者帳號 (commit 1)
```zsh
// API
http://localhost:5000/list-all-users
```
### POSTMAN DEMO  (commit 1)
![image](https://upload.cc/i1/2022/04/19/7QZwDW.png)
<br></br>

## 列出所有使用者詳細資料 (commit 3)
```zsh
// API
http://localhost:5000/get-user-detailed-information
```
### POSTMAN DEMO  (commit 3)
![image](https://upload.cc/i1/2022/04/19/cHpqP2.png)
<br></br>
