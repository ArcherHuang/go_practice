
- 第一週回家作業，實作主題：霹靂布袋戲角色資料庫系統
    - 事前準備        
        - 使用事先建立好的 Go 專案範本開發 `版本_1`
        - 可使用 [Swagger Editor](https://editor.swagger.io/) 匯入提供的 swagger.yaml 規則，查看 API 規格
            - 或者使用 Postman 匯入「布袋戲資料庫.postman_collection.json」
    - 實作目標：利用 `gin` 實作五個 RESTful APIs
        - 取得全部資料 [GET] http://localhost:8080/role
        - 新增單筆資料 [POST] http://localhost:8080/role
        - 取得單筆資料 [GET] http://localhost:8080/role/:id
        - 更新單筆資料 [PUT] http://localhost:8080/role/:id
        - 刪除單筆資料 [DELETE] http://localhost:8080/role/:id
    - 實作方式
        - 利用 `slice` 管理資料集合（不透過資料庫）**(使用 `版本_1` Go 專案範本)**

